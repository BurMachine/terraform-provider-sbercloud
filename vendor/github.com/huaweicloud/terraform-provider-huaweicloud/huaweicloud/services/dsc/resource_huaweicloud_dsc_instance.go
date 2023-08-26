// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product DSC
// ---------------------------------------------------------------

package dsc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chnsz/golangsdk"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/jmespath/go-jmespath"
)

const (
	cloudServiceType = "hws.service.type.sdg"

	resourceTypeBase = "hws.resource.type.dsc.base"
	resourceTypeDB   = "hws.resource.type.dsc.db"
	resourceTypeObs  = "hws.resource.type.dsc.obs"

	resourceSpecCodeProBase      = "base_professional"
	resourceSpecCodeStandardBase = "base_standard"

	resourceSpecCodeProDB      = "DB_professional"
	resourceSpecCodeStandardDB = "DB_standard"

	resourceSpecCodeProObs      = "OBS_professional"
	resourceSpecCodeStandardObs = "OBS_standard"

	dscBaseProductIdStandard = "OFFI571544772268847116"
	dscBaseProductIdPro      = "OFFI571544772268847108"
	dscObsProductIdStandard  = "OFFI571545095458668555"
	dscObsProductIdPro       = "OFFI571545095458668547"
	dscDBProductIdStandard   = "OFFI571544962480533516"
	dscDBProductIdPro        = "OFFI571544962480533508"

	resourceSizeMeasureIdObs = 47
	resourceSizeMeasureIdDB  = 30
)

func ResourceDscInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDscInstanceCreate,
		ReadContext:   resourceDscInstanceRead,
		DeleteContext: resourceDscInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"edition": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The edition of DSC.`,
				ValidateFunc: validation.StringInSlice([]string{
					resourceSpecCodeStandardBase, resourceSpecCodeProBase,
				}, false),
			},
			"charging_mode": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Billing mode.`,
				ValidateFunc: validation.StringInSlice([]string{
					"prePaid",
				}, false),
			},
			"period_unit": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The charging period unit.`,
				ValidateFunc: validation.StringInSlice([]string{
					"month", "year",
				}, false),
			},
			"period": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				Description:  `The charging period.`,
				ValidateFunc: validation.IntBetween(1, 9),
			},
			"auto_renew": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Whether auto renew is enabled. Valid values are "true" and "false".`,
				ValidateFunc: validation.StringInSlice([]string{
					"true", "false",
				}, false),
			},
			"obs_expansion_package": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Purchase OBS expansion packages. One OBS expansion package offers 1 TB of OBS storage.`,
			},
			"database_expansion_package": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Purchase database expansion packages. One expansion package offers one database.`,
			},
		},
	}
}

func resourceDscInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	// createDscInstance: create a DSC.
	var (
		createDscInstanceHttpUrl = "v1/{project_id}/period/order"
		createDscInstanceProduct = "dsc"
	)
	createDscInstanceClient, err := config.NewServiceClient(createDscInstanceProduct, region)
	if err != nil {
		return diag.Errorf("error creating DscInstance Client: %s", err)
	}

	createDscInstancePath := createDscInstanceClient.Endpoint + createDscInstanceHttpUrl
	createDscInstancePath = strings.ReplaceAll(createDscInstancePath, "{project_id}", createDscInstanceClient.ProjectID)

	createDscInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	createDscInstanceOpt.JSONBody = utils.RemoveNil(buildCreateDscInstanceBodyParams(d, config))
	createDscInstanceResp, err := createDscInstanceClient.Request("POST", createDscInstancePath, &createDscInstanceOpt)
	if err != nil {
		return diag.Errorf("error creating DscInstance: %s", err)
	}

	createDscInstanceRespBody, err := utils.FlattenResponse(createDscInstanceResp)
	if err != nil {
		return diag.FromErr(err)
	}

	orderId, err := jmespath.Search("order_id", createDscInstanceRespBody)
	if err != nil {
		return diag.Errorf("error creating DscInstance: ID is not found in API response")
	}

	// auto pay
	var (
		payOrderHttpUrl = "v3/orders/customer-orders/pay"
		payOrderProduct = "bss"
	)
	payOrderClient, err := config.NewServiceClient(payOrderProduct, region)
	if err != nil {
		return diag.Errorf("error creating BSS Client: %s", err)
	}

	payOrderPath := payOrderClient.Endpoint + payOrderHttpUrl

	payOrderOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			204,
		},
	}
	payOrderOpt.JSONBody = utils.RemoveNil(buildPayOrderBodyParams(orderId.(string)))
	_, err = payOrderClient.Request("POST", payOrderPath, &payOrderOpt)
	if err != nil {
		return diag.Errorf("error pay order=%s: %s", d.Id(), err)
	}

	bssClient, err := config.BssV2Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating BSS v2 client: %s", err)
	}
	err = common.WaitOrderComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	resourceId, err := common.WaitOrderResourceComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resourceId)
	return resourceDscInstanceRead(ctx, d, meta)
}

func buildCreateDscInstanceBodyParams(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"regionId":         config.GetRegion(d),
		"cloudServiceType": cloudServiceType,
		"periodNum":        utils.ValueIngoreEmpty(d.Get("period")),
		"productInfos":     buildCreateDscInstanceRequestBodyProductInfos(d),
	}

	chargingMode := d.Get("charging_mode").(string)
	if chargingMode == "prePaid" {
		bodyParams["chargingMode"] = 0
	}

	periodUnit := d.Get("period_unit").(string)
	if periodUnit == "month" {
		bodyParams["periodType"] = 2
	} else {
		bodyParams["periodType"] = 3
	}

	autoRenew := d.Get("auto_renew").(string)
	if autoRenew == "true" {
		bodyParams["isAutoRenew"] = 1
	} else {
		bodyParams["isAutoRenew"] = 0
	}

	return bodyParams
}

func buildCreateDscInstanceRequestBodyProductInfos(d *schema.ResourceData) []map[string]interface{} {
	rst := make([]map[string]interface{}, 0, 3)
	edition := d.Get("edition").(string)

	if edition == resourceSpecCodeStandardBase {
		rst = append(rst, map[string]interface{}{
			"cloudServiceType": cloudServiceType,
			"productId":        dscBaseProductIdStandard,
			"resourceType":     resourceTypeBase,
			"resourceSpecCode": resourceSpecCodeStandardBase,
		})

		if size, ok := d.GetOk("obs_expansion_package"); ok {
			rst = append(rst, map[string]interface{}{
				"cloudServiceType":      cloudServiceType,
				"productId":             dscObsProductIdStandard,
				"resourceType":          resourceTypeObs,
				"resourceSpecCode":      resourceSpecCodeStandardObs,
				"resourceSize":          utils.ValueIngoreEmpty(size),
				"resourceSizeMeasureId": resourceSizeMeasureIdObs,
			})
		}

		if size, ok := d.GetOk("database_expansion_package"); ok {
			rst = append(rst, map[string]interface{}{
				"cloudServiceType":      cloudServiceType,
				"productId":             dscDBProductIdStandard,
				"resourceType":          resourceTypeDB,
				"resourceSpecCode":      resourceSpecCodeStandardDB,
				"resourceSize":          utils.ValueIngoreEmpty(size),
				"resourceSizeMeasureId": resourceSizeMeasureIdDB,
			})
		}
	} else {
		rst = append(rst, map[string]interface{}{
			"cloudServiceType": cloudServiceType,
			"productId":        dscBaseProductIdPro,
			"resourceType":     resourceTypeBase,
			"resourceSpecCode": resourceSpecCodeProBase,
		})

		if size, ok := d.GetOk("obs_expansion_package"); ok {
			rst = append(rst, map[string]interface{}{
				"cloudServiceType":      cloudServiceType,
				"productId":             dscObsProductIdPro,
				"resourceType":          resourceTypeObs,
				"resourceSpecCode":      resourceSpecCodeProObs,
				"resourceSize":          utils.ValueIngoreEmpty(size),
				"resourceSizeMeasureId": resourceSizeMeasureIdObs,
			})
		}

		if size, ok := d.GetOk("database_expansion_package"); ok {
			rst = append(rst, map[string]interface{}{
				"cloudServiceType":      cloudServiceType,
				"productId":             dscDBProductIdPro,
				"resourceType":          resourceTypeDB,
				"resourceSpecCode":      resourceSpecCodeProDB,
				"resourceSize":          utils.ValueIngoreEmpty(size),
				"resourceSizeMeasureId": resourceSizeMeasureIdDB,
			})
		}
	}

	return rst
}

func buildPayOrderBodyParams(orderId string) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"order_id":     orderId,
		"use_coupon":   "NO",
		"use_discount": "NO",
	}
	return bodyParams
}

func resourceDscInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	var mErr *multierror.Error

	// getDscInstance: Query the DSC instance
	var (
		getDscInstanceHttpUrl = "v1/{project_id}/period/product/specification"
		getDscInstanceProduct = "dsc"
	)
	getDscInstanceClient, err := config.NewServiceClient(getDscInstanceProduct, region)
	if err != nil {
		return diag.Errorf("error creating DscInstance Client: %s", err)
	}

	getDscInstancePath := getDscInstanceClient.Endpoint + getDscInstanceHttpUrl
	getDscInstancePath = strings.ReplaceAll(getDscInstancePath, "{project_id}", getDscInstanceClient.ProjectID)

	getDscInstanceOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
	}
	getDscInstanceResp, err := getDscInstanceClient.Request("GET", getDscInstancePath, &getDscInstanceOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving DscInstance")
	}

	getDscInstanceRespBody, err := utils.FlattenResponse(getDscInstanceResp)
	if err != nil {
		return diag.FromErr(err)
	}

	dscOrder, err := jmespath.Search("orderInfo[?productInfo.resourceType=='hws.resource.type.dsc.base']",
		getDscInstanceRespBody)
	if err != nil {
		return diag.Errorf("error getting the instance info: base info not found in API response")
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("period_unit", parsePeriodUnit(utils.PathSearch("[0].periodType", dscOrder, nil))),
		d.Set("period", utils.PathSearch("[0].periodNum", dscOrder, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func parsePeriodUnit(periodType interface{}) string {
	pUnit := fmt.Sprintf("%v", periodType)
	if pUnit == "3" {
		return "year"

	} else {
		return "month"
	}
}

func resourceDscInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)

	if err := common.UnsubscribePrePaidResource(d, config, []string{d.Id()}); err != nil {
		return diag.Errorf("Error unsubscribing DSC order = %s: %s", d.Id(), err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForBmsInstanceDelete(ctx, d, meta),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        20 * time.Second,
		PollInterval: 10 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error deleting DSC instance: %s", err)
	}

	return nil
}

func waitForBmsInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	config := meta.(*config.Config)
	region := config.GetRegion(d)

	return func() (interface{}, string, error) {
		// getDscInstance: Query the DSC instance
		var (
			getDscInstanceHttpUrl = "v1/{project_id}/period/product/specification"
			getDscInstanceProduct = "dsc"
		)
		getDscInstanceClient, err := config.NewServiceClient(getDscInstanceProduct, region)
		if err != nil {
			return nil, "error", fmt.Errorf("error creating DscInstance Client: %s", err)
		}

		getDscInstancePath := getDscInstanceClient.Endpoint + getDscInstanceHttpUrl
		getDscInstancePath = strings.ReplaceAll(getDscInstancePath, "{project_id}", getDscInstanceClient.ProjectID)

		getDscInstanceOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
		}
		getDscInstanceResp, err := getDscInstanceClient.Request("GET", getDscInstancePath, &getDscInstanceOpt)

		if err != nil {
			return nil, "error", fmt.Errorf("error retrieving DscInstance: %s", err)
		}

		getDscInstanceRespBody, err := utils.FlattenResponse(getDscInstanceResp)
		if err != nil {
			return nil, "error", fmt.Errorf("error retrieving DscInstance: %s", err)
		}

		orderInfo := utils.PathSearch("orderInfo", getDscInstanceRespBody, []interface{}{})
		orders := orderInfo.([]interface{})
		if len(orders) == 0 {
			return orders, "COMPLETE", nil
		} else {
			return nil, "PENDING", nil
		}
	}
}