package cce

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk/openstack/cce/v3/partitions"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func ResourcePartition() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePartitionCreate,
		ReadContext:   resourcePartitionRead,
		UpdateContext: resourcePartitionUpdate,
		DeleteContext: resourcePartitionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePartitionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partition_subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"container_subnet_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func buildPartitionSubnetID(d *schema.ResourceData) partitions.HostNetwork {
	return partitions.HostNetwork{
		SubnetID: d.Get("partition_subnet_id").(string),
	}
}

func buildContainerSubnetIDs(d *schema.ResourceData) []partitions.ContainerNetwork {
	networkRaw := d.Get("container_subnet_ids").(*schema.Set)
	containerNetwork := make([]partitions.ContainerNetwork, networkRaw.Len())
	for i, raw := range networkRaw.List() {
		containerNetwork[i] = partitions.ContainerNetwork{
			SubnetID: raw.(string),
		}
	}
	return containerNetwork
}

func resourcePartitionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("Error creating CCE Partition client: %s", err)
	}

	// wait for the cce cluster to become available
	clusterid := d.Get("cluster_id").(string)
	stateCluster := &resource.StateChangeConf{
		Target:       []string{"Available"},
		Refresh:      waitForClusterAvailable(cceClient, clusterid),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 5 * time.Second,
	}
	_, err = stateCluster.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error waiting for CCE cluster to be Available: %s", err)
	}

	createOpts := partitions.CreateOpts{
		Kind:       "Partition",
		ApiVersion: "v3",
		Metadata: partitions.CreateMetaData{
			Name: d.Get("availability_zone").(string),
		},
		Spec: partitions.Spec{
			Category:          d.Get("category").(string),
			PublicBorderGroup: d.Get("availability_zone").(string),
			HostNetwork:       buildPartitionSubnetID(d),
			ContainerNetwork:  buildContainerSubnetIDs(d),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	s, err := partitions.Create(cceClient, clusterid, createOpts).Extract()
	if err != nil {
		return diag.Errorf("Error creating Partition: %s", err)
	}
	d.SetId(s.Metadata.Name)

	return resourcePartitionRead(ctx, d, meta)
}

func resourcePartitionRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("Error creating CCE client: %s", err)
	}
	clusterid := d.Get("cluster_id").(string)
	s, err := partitions.Get(cceClient, clusterid, d.Id()).Extract()

	if err != nil {
		return common.CheckDeletedDiag(d, err, "Error retrieving CCE Partition")
	}

	mErr := multierror.Append(nil,
		d.Set("region", cfg.GetRegion(d)),
		d.Set("name", s.Metadata.Name),
		d.Set("category", s.Spec.Category),
		d.Set("availability_zone", s.Spec.PublicBorderGroup),
		d.Set("partition_subnet_id", s.Spec.HostNetwork.SubnetID),
	)

	var containerNetworkIDs []string
	for _, network := range s.Spec.ContainerNetwork {
		containerNetworkIDs = append(containerNetworkIDs, network.SubnetID)
	}
	mErr = multierror.Append(mErr, d.Set("container_subnet_ids", containerNetworkIDs))

	if err = mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("Error setting CCE Partition fields: %s", err)
	}
	return nil
}

func resourcePartitionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	cceClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("Error creating CCE client: %s", err)
	}

	if d.HasChange("container_subnet_ids") {
		var updateOpts partitions.UpdateOpts
		updateOpts.Metadata.ContainerNetwork = buildContainerSubnetIDs(d)

		clusterid := d.Get("cluster_id").(string)
		_, err = partitions.Update(cceClient, clusterid, d.Id(), updateOpts).Extract()
		if err != nil {
			return diag.Errorf("Error updating cce partition: %s", err)
		}
	}

	return resourceNodeRead(ctx, d, meta)
}

func resourcePartitionDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("Error creating CCE client: %s", err)
	}

	clusterid := d.Get("cluster_id").(string)
	err = partitions.Delete(cceClient, clusterid, d.Id()).ExtractErr()
	if err != nil {
		return diag.Errorf("Error deleting CCE partition: %s", err)
	}

	return nil
}

func resourcePartitionImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		err := fmt.Errorf("invalid format specified for CCE Partition. Format must be <cluster id>/<partition name>")
		return nil, err
	}

	clusterID := parts[0]
	partitionName := parts[1]

	d.SetId(partitionName)
	d.Set("cluster_id", clusterID)

	return []*schema.ResourceData{d}, nil
}
