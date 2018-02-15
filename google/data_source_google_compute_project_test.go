package google

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceGoogleComputeProject_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.google_compute_project.project"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleComputeProject_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "project_number"),
					resource.TestCheckResourceAttrSet(resourceName, "default_service_account"),
					resource.TestCheckResourceAttrSet(resourceName, "self_link"),
				),
			},
		},
	})
}

const testAccCheckGoogleComputeProject_basic = `
data "google_compute_project" "project" { }
`
