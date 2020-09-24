---
subcategory: ""
layout: ""
page_title: "terraform-provider-circleci: circleci_project"
description: |-
  circleci_project manages projects in CircleCI.
---

# Resource: `circleci_project`

circleci_project manages projects in CircleCI.

## Example Usage

```terraform
resource "circleci_project" "example" {
    name     = "MyCircleProject"
    env_vars {
      SOME_VARIABLE = "MyVariableValue"
    }
}
```

## Schema

### Required

- **name** (String, Required) The name of the CircleCI project.

### Optional

- **env_vars** (Map, Optional) Map of Environment variables for the project 

