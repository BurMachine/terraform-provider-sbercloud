package dli

import (
	"context"
	"regexp"
	"time"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/dli/v2/batches"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
)

func ResourceDliSparkJobV2() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceDliSparkJobV2Create,
		ReadContext:   ResourceDliSparkJobV2Read,
		DeleteContext: ResourceDliSparkJobV2Delete,

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"queue_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_parameters": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"main_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"jars": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"python_files": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"files": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dependent_packages": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`[\w-.]/*`), "Only digits, letters,"+
								" dots (.), underscores (_), and hyphens (-) are allowed for group name."),
						},
						"packages": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										ValidateFunc: validation.StringInSlice([]string{
											jarFile, pythonFile, userFile}, false),
									},
									"package_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"configurations": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"modules": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"specification": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"A", "B", "C",
				}, false),
			},
			"executor_memory": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"executor_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"executors": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"driver_memory": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"driver_cores": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func buildDliSaprkGroups(packages []interface{}) []batches.Group {
	result := make([]batches.Group, len(packages))
	for i, val1 := range packages {
		resources := val1.(map[string]interface{})
		group := batches.Group{
			Name: resources["group_name"].(string),
		}

		apps := resources["packages"].([]interface{})
		res := make([]batches.Resource, len(apps))
		for j, val2 := range apps {
			app := val2.(map[string]interface{})
			res[j] = batches.Resource{
				Type: app["type"].(string),
				Name: app["package_name"].(string),
			}
		}
		group.Resources = res

		result[i] = group
	}

	return result
}

func buildDliSaprkJobCreateOpts(d *schema.ResourceData) batches.CreateOpts {
	result := batches.CreateOpts{
		Queue: d.Get("queue_name").(string),
		Name:  d.Get("name").(string),
		File:  d.Get("app_name").(string),
		// This parameter is required according to API ducumentation.
		ClassName:      golangsdk.MaybeString(d.Get("main_class").(string)),
		Groups:         buildDliSaprkGroups(d.Get("dependent_packages").([]interface{})),
		Configurations: d.Get("configurations").(map[string]interface{}),
		ExecutorMemory: d.Get("executor_memory").(string),
		ExecutorCores:  d.Get("executor_cores").(int),
		NumExecutors:   d.Get("executors").(int),
		DriverMemory:   d.Get("driver_memory").(string),
		DriverCores:    d.Get("driver_cores").(int),
		MaxRetryTimes:  d.Get("max_retries").(int),
		Specification:  d.Get("specification").(string),
	}
	if params, ok := d.GetOk("app_parameters"); ok {
		result.Arguments = utils.ExpandToStringList(params.([]interface{}))
	}
	if jars, ok := d.GetOk("jars"); ok {
		result.Jars = utils.ExpandToStringList(jars.([]interface{}))
	}
	if pyFiles, ok := d.GetOk("python_files"); ok {
		result.PythonFiles = utils.ExpandToStringList(pyFiles.([]interface{}))
	}
	if files, ok := d.GetOk("files"); ok {
		result.Files = utils.ExpandToStringList(files.([]interface{}))
	}
	if modules, ok := d.GetOk("modules"); ok {
		result.Modules = utils.ExpandToStringList(modules.([]interface{}))
	}

	return result
}

func ResourceDliSparkJobV2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	c, err := config.DliV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud DLI v2 client: %s", err)
	}

	resp, err := batches.Create(c, buildDliSaprkJobCreateOpts(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating spark job: %s", err)
	}

	d.SetId(resp.ID)

	return ResourceDliSparkJobV2Read(ctx, d, meta)
}

func setDliSparkJobParameters(d *schema.ResourceData, resp *batches.CreateResp) error {
	mErr := multierror.Append(nil,
		d.Set("queue_name", resp.Queue),
		d.Set("name", resp.Name),
		d.Set("created_at", time.Unix(int64(resp.CreateTime)/1000, 0).Format("2006-01-02 15:04:05")),
		d.Set("owner", resp.Owner),
	)
	if mErr.ErrorOrNil() != nil {
		return mErr
	}
	return nil
}

func ResourceDliSparkJobV2Read(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	c, err := config.DliV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud DLI v2 client: %s", err)
	}

	resp, err := batches.Get(c, d.Id())
	if err != nil {
		return fmtp.DiagErrorf("Error getting spark job: %s", err)
	}

	err = setDliSparkJobParameters(d, resp)
	if err != nil {
		return fmtp.DiagErrorf("An error occurred during spark job resource parameter setting: %s", err)
	}
	return nil
}

func ResourceDliSparkJobV2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	c, err := config.DliV2Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud DLI v2 client: %s", err)
	}

	resp, err := batches.GetState(c, d.Id())
	if err != nil {
		return fmtp.DiagErrorf("Error getting spark job status: %s", err)
	}

	switch resp.State {
	// The spark job can be cancel while status is 'starting', 'running' or 'recovering'.
	case batches.StateStarting, batches.StateRunning, batches.StateRecovering:
		err = batches.Delete(c, d.Id()).ExtractErr()
		if err != nil {
			return fmtp.DiagErrorf("Unable to cancel spark job: %s", err)
		}
	}

	err = checkDliSparkJobCancelResult(ctx, c, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func checkDliSparkJobCancelResult(ctx context.Context, client *golangsdk.ServiceClient, jobId string,
	timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"Pending"},
		Target:  []string{"Success"},
		Refresh: func() (interface{}, string, error) {
			resp, err := batches.GetState(client, jobId)
			if err == nil && (resp.State == batches.StateDead ||
				resp.State == batches.StateSuccess) {
				return true, "Success", nil
			}
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault400); ok {
					return true, "Success", nil
				}
				return nil, "", nil
			}
			return true, "Pending", nil
		},
		Timeout:      timeout,
		PollInterval: 10 * timeout,
		Delay:        20 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmtp.Errorf("error waiting for DLI spark job (%s) to be cancelled: %s", jobId, err)
	}
	return nil
}
