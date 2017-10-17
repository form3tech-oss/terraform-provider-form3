---
layout: "form3"
page_title: "Form3: form3_user"
sidebar_current: "docs-form3-resource-user"
description: |-
  Provides a Form3 user resource.
---

# form3\_user

A [user](http://api-docs.form3.tech/#user-resource) resource.

## Example Usage

```hcl
# Add a user to your form3 account
resource "form3_user" "main" {
  name      = "a-user"
  team_uuid = "870ed937-bc6e-4d8b-a9a5-d7f9f2412fa3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (String, Required) The name of this user.
* `team_uuid` - (String, Required) Unique identifier for the team this user
  is being created for.

## Attributes Reference

The following attributes are exported:

* `name` - The name of this user.
* `team_uuid` - Unique identifier for the team this user belongs to.