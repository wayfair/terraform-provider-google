package google

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGoogleComputeProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeProjectRead,
		Schema: map[string]*schema.Schema{
			"default_service_account": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGoogleComputeProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	projectCompResource, err := config.clientCompute.Projects.Get(project).Do()
	if err != nil {
		return fmt.Errorf("Error reading project resource: %s", err)
	}

	projectResManager, err := config.clientResourceManager.Projects.Get(project).Do()
	if err != nil {
		return fmt.Errorf("Error reading project resource: %s", err)
	}

	d.SetId(projectResManager.ProjectId)
	d.Set("project_number", strconv.FormatInt(projectResManager.ProjectNumber, 10))
	d.Set("default_service_account", projectCompResource.DefaultServiceAccount)
	d.Set("self_link", projectCompResource.SelfLink)

	return nil
}
