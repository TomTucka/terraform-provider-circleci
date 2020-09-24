package circleci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEnvironmentVariable(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		// TODO: CheckDestroy: ,
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentVariableConfig("testev", "secret_test", "secret_value"),
			},
			importStep("circleci_environment_variable.test"),
			{
				Config: testAccEnvironmentVariableConfig("testev", "secret_test", "secret_value"),
			},
			importStep("circleci_environment_variable.test"),
		},
	})
}

func testAccEnvironmentVariableConfig(project, name, value string) string {
	return fmt.Sprintf(`
resource "circleci_environment_variable" "%[2]s" {
  project = "%[1]s"
  name    = "%[2]s"
  value   = "%[3]s"
}`, project, name, value)
}
