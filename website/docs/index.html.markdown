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
  token        = var.circleci_token 
  organization = "YourOrganizationName" 
}
```