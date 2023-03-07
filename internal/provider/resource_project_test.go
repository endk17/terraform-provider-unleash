package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/philips-labs/terraform-provider-unleash/utils"
)

func TestAccResourceProject(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("unleash_project.foo", "project_id", regexp.MustCompile("^bar")),
					resource.TestMatchResourceAttr("unleash_project.foo", "name", regexp.MustCompile("^bar")),
					resource.TestMatchResourceAttr("unleash_project.foo", "description", regexp.MustCompile("^xyz")),
				),
			},
		},
	})
}

var TestAccResourceProjects = fmt.Sprintf(`
resource "unleash_project" "foo" {
  project_id 	= "bar"
  name			= "bar"
  description 	= "xyz_%s"
}
`, utils.RandomString(4))
