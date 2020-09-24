---
subcategory: ""
layout: ""
page_title: "terraform-provider-circleci: circleci_environment_variable"
description: |-
  circleci_environment_variable manages environment variables on CircleCI projects.
---

# Resource: `circleci_environment_variable`

  circleci_environment_variable manages environment variables on CircleCI projects.

## Example Usage

```terraform
resource "circleci_environment_variable" "example" {
    project    = "MyCircleProject"
    name       = "SOME_VARIABLE"
    value      = "MyVariableValue"
}
```

## Schema

### Required

- **project** (String, Required) The name of the CircleCI project.
- **name** (String, Required) The name of the environment variable.
- **value** (String, Required) The value of the environment variable.

