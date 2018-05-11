# terraform-provider-circleci

[Terraform][] provider for [CircleCI][].

## Usage

### Provider

The CircleCI provider is used to interact with CircleCI resources.

The provider allows you to manage your CircleCI projects and their environment
variables easily. It needs to be configured with the proper credentials before
it can be used.

#### Example Usage

```hdl
# Configure the CircleCI Provider
provider "circleci" {
  token        = "${var.circleci_token}"
  organization = "${var.circleci_organization}"
}
```

#### Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Optional) This is the CircleCI API token. It must be provided,
  but it can also be sourced from the `CIRCLECI_TOKEN` environment variable.

* `organization` - (Optional) This is the organization/account to be managed.
It must be provided, but it can also be sourced from the `CIRCLECI_ORGANIZATION`
environment variable.

### circleci_project

Provides a CircleCI project resource.

This resource allows you to start/stop building projects from your organization.
When applied, a project is enabled. When destroyed, that project will be disabled.

#### Example Usage

```hdl
# Start building a project on CircleCI
resource "circleci_project" "myproj" {
  name = "myproj"
}
```

#### Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the project/repository.

#### Import

CircleCI projects can be imported using the name, e.g.

```sh
$ terraform import circleci_project.myproj myproj
```

### circleci_envvar

Provides a CircleCI envvar resource.

This resource allows you to add/update/remove environment variables from your
projects. When applied, an environment variable is created. When destroyed, that
environment variable will be removed.

#### Example Usage

```hdl
# Add an environment variable to a project
resource "circleci_envvar" "myproj-some-token" {
  project = "${circleci_project.myproj.name}"
  name    = "SOME_TOKEN"
  value   = "a1b2c3d4e5f6g7h8i9j0"
}
```

#### Argument Reference

The following arguments are supported:

* `project` - (Required) The project in which the env var will be added.

* `name` - (Required) The environment variable name.

* `value` - (Required) The environment variable value.

#### Import

CircleCI environment variables can be imported by using the project + the variable
name, e.g.

```sh
$ terraform import circleci_envvar.myproj-some-token myproj-SOME_TOKEN
```

[Terraform]: https://www.terraform.io
[CircleCI]: https://circleci.com
