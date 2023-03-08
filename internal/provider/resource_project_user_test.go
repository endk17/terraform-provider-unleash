package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/philips-labs/terraform-provider-unleash/utils"
)

func TestAccResourceProjectUsers(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("unleash_project_user.foo", "user_id", regexp.MustCompile("^123")),
					resource.TestMatchResourceAttr("unleash_project_user.foo", "project_id", regexp.MustCompile("^bar")),
					resource.TestMatchResourceAttr("unleash_project_user.foo", "role_id", regexp.MustCompile("^456")),
				),
			},
		},
	})
}

var TestAccResourceProjectUser = fmt.Sprintf(`
resource "unleash_project_user" "foo" {
  user_id 	    = 123
  project_id	= "bar_%s"
  role_id 	    = 456
}
`, utils.RandomString(4))
