---
layout: ""
page_title: "Provider: CircleCI"
description: |-
  The CircleCI provider provides resources to interact with and manage CircleCI projects.
---

# CircleCI Provider

The CircleCI provider provides resources to interact with and manage CircleCI projects.

## Example Usage

```terraform
provider "circleci" {
  token        = var.circleci_token # optionally use CIRCLECI_TOKEN env var
  organization = "YourOrganizationName" # optionally use CIRCLECI_ORGANIZATION env var
  vcs_type     = "github" # optionally use CIRCLECI_VCS_TYPE env var
}
```

## Schema

- **token** (String, Required) CircleCI API Token. You can create this in your account dashboard. Can be specified with the CIRCLECI_TOKEN environment variable.

- **organization** (String, Required) Configure what organization you are using in CircleCI. Can be specified with the CIRCLECI_ORGANIZATION environment variable.

- **vcs_type** (String, Required) Configure what version control system type your project uses. You may currently select either ‘github’ or ‘bitbucket’. Can be specified with the CIRCLECI_VCS_TYPE environment variable.