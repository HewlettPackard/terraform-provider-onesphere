---
layout: "onesphere"
page_title: "onesphere: user"
sidebar_current: "docs-onesphere-user"
description: |-
  Creates an user.
---

# onesphere\_user\

Creates an user.

## Example Usage

```js
resource "onesphere_user" "default" {
  name     = "user_name"
  email    = "user_email"
  password = "user_password"
  role     = "administrator|analyst|consumer|project-creator"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique user name for the resource.

* `email` - (Required) A unique user email for the resource.

* `password` - (Required) A unique user password for the resource.

* `role` - (Required) A user role can be administrator|analyst|consumer|project-creator for the resource.

---