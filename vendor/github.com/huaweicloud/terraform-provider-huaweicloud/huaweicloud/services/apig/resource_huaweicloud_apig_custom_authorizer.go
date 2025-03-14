package apig

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk/openstack/apigw/dedicated/v2/authorizers"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

type AuthType string

const (
	AuthTypeFrontend AuthType = "FRONTEND"
	AuthTypeBackend  AuthType = "BACKEND"
)

func ResourceApigCustomAuthorizerV2() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomAuthorizerCreate,
		ReadContext:   resourceCustomAuthorizerRead,
		UpdateContext: resourceCustomAuthorizerUpdate,
		DeleteContext: resourceCustomAuthorizerDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceCustomAuthorizerImportState,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The region where the custom authorizer is located.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the dedicated instance to which the custom authorizer belongs.",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringMatch(regexp.MustCompile(`^[A-Za-z][\w]*$`),
						"Only letters, digits and underscores (_) are allowed, and must start with a letter."),
					validation.StringLenBetween(3, 64),
				),
				Description: "The name of the custom authorizer.",
			},
			"function_urn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URN of the FGS function.",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  string(AuthTypeFrontend),
				ValidateFunc: validation.StringInSlice([]string{
					string(AuthTypeFrontend), string(AuthTypeBackend),
				}, false),
				Description: "The custom authorization type",
			},
			"is_body_send": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to send the body.",
			},
			"cache_age": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 3600),
				Description:  "The maximum cache age.",
			},
			"user_data": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 2048),
				Description:  "The user data for custom authorizer function.",
			},
			// The parameter identity only required if type is 'FRONTEND'.
			"identity": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the parameter to be verified.",
						},
						"location": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"HEADER", "QUERY",
							}, false),
							Description: "The parameter location.",
						},
						"validation": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(1, 2048),
							Description:  "The parameter verification expression.",
						},
					},
				},
				Description: "The array of one or more parameter identities of the custom authorizer.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the custom authorizer.",
			},
		},
	}
}

func buildIdentities(identities *schema.Set) []authorizers.AuthCreateIdentitiesReq {
	if identities.Len() < 1 {
		return nil
	}

	result := make([]authorizers.AuthCreateIdentitiesReq, identities.Len())
	for i, val := range identities.List() {
		identity := val.(map[string]interface{})
		validContent := identity["validation"].(string)
		result[i] = authorizers.AuthCreateIdentitiesReq{
			Name:       identity["name"].(string),
			Location:   identity["location"].(string),
			Validation: &validContent,
		}
	}
	return result
}

func buildCustomAuthorizerOpts(d *schema.ResourceData) (authorizers.CustomAuthOpts, error) {
	var (
		t          = d.Get("type").(string) // The 'authType' is easily confused with 'AuthorizerType', and 'type' is a keyword.
		identities = d.Get("identity").(*schema.Set)
	)
	if identities.Len() > 0 && t != string(AuthTypeFrontend) {
		return authorizers.CustomAuthOpts{}, fmt.Errorf("the identities can only be set when the type is 'FRONTEND'")
	}

	return authorizers.CustomAuthOpts{
		Name:           d.Get("name").(string),
		Type:           t,
		AuthorizerType: "FUNC", // The custom authorizer only support 'FUNC'.
		AuthorizerURI:  d.Get("function_urn").(string),
		IsBodySend:     utils.Bool(d.Get("is_body_send").(bool)),
		TTL:            utils.Int(d.Get("cache_age").(int)),
		UserData:       utils.String(d.Get("user_data").(string)),
		Identities:     buildIdentities(identities),
	}, nil
}

func resourceCustomAuthorizerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	client, err := cfg.ApigV2Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating APIG v2 client: %s", err)
	}

	opt, err := buildCustomAuthorizerOpts(d)
	if err != nil {
		return diag.FromErr(err)
	}
	instanceId := d.Get("instance_id").(string)
	resp, err := authorizers.Create(client, instanceId, opt).Extract()
	if err != nil {
		return common.CheckDeletedDiag(d, err, "custom authorizer")
	}
	d.SetId(resp.ID)

	return resourceCustomAuthorizerRead(ctx, d, meta)
}

func flattenCustomAuthorizerIdentities(identities []authorizers.Identity) []map[string]interface{} {
	if len(identities) < 1 {
		return nil
	}

	result := make([]map[string]interface{}, len(identities))
	for i, val := range identities {
		result[i] = map[string]interface{}{
			"name":       val.Name,
			"location":   val.Location,
			"validation": val.Validation,
		}
	}
	return result
}

func resourceCustomAuthorizerRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg        = meta.(*config.Config)
		region     = cfg.GetRegion(d)
		instanceId = d.Get("instance_id").(string)
	)
	client, err := cfg.ApigV2Client(region)
	if err != nil {
		return diag.Errorf("error creating APIG v2 client: %s", err)
	}
	resp, err := authorizers.Get(client, instanceId, d.Id()).Extract()
	if err != nil {
		return common.CheckDeletedDiag(d, err, fmt.Sprintf("unable to get the custom authorizer (%s) form server", d.Id()))
	}
	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("name", resp.Name),
		d.Set("function_urn", resp.AuthorizerURI),
		d.Set("type", resp.Type),
		d.Set("is_body_send", resp.IsBodySend),
		d.Set("cache_age", resp.TTL),
		d.Set("user_data", resp.UserData),
		d.Set("created_at", resp.CreateTime),
		d.Set("identity", flattenCustomAuthorizerIdentities(resp.Identities)),
	)
	if err = mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error saving custom authorizer fields: %s", err)
	}
	return nil
}

func resourceCustomAuthorizerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg = meta.(*config.Config)

		instanceId   = d.Get("instance_id").(string)
		authorizerId = d.Id()
	)
	client, err := cfg.ApigV2Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating APIG v2 client: %s", err)
	}
	opt, err := buildCustomAuthorizerOpts(d)
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = authorizers.Update(client, instanceId, authorizerId, opt).Extract()
	if err != nil {
		return diag.Errorf("error updating APIG custom authorizer (%s): %s", authorizerId, err)
	}

	return resourceCustomAuthorizerRead(ctx, d, meta)
}

func resourceCustomAuthorizerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg = meta.(*config.Config)

		instanceId   = d.Get("instance_id").(string)
		authorizerId = d.Id()
	)
	client, err := cfg.ApigV2Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating APIG v2 client: %s", err)
	}

	err = authorizers.Delete(client, instanceId, authorizerId).ExtractErr()
	if err != nil {
		return diag.Errorf("error deleting custom authorizer (%s) from the instance (%s): %s",
			authorizerId, instanceId, err)
	}
	return nil
}

// The ID cannot find on the console, so we need to import by authorizer name.
func resourceCustomAuthorizerImportState(_ context.Context, d *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid format specified for import ID, must be <instance_id>/<name>")
	}
	var (
		cfg        = meta.(*config.Config)
		instanceId = parts[0]
		name       = parts[1]

		opt = authorizers.ListOpts{
			Name: name,
		}
	)
	client, err := cfg.ApigV2Client(cfg.GetRegion(d))
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf("error creating APIG v2 client: %s", err)
	}
	pages, err := authorizers.List(client, instanceId, opt).AllPages()
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf("error retrieving custom authorizer: %s", err)
	}

	resp, err := authorizers.ExtractCustomAuthorizers(pages)
	if len(resp) < 1 {
		return []*schema.ResourceData{d}, fmt.Errorf("unable to find the custom authorizer (%s) form server: %s",
			name, err)
	}
	d.SetId(resp[0].ID)

	return []*schema.ResourceData{d}, d.Set("instance_id", instanceId)
}
