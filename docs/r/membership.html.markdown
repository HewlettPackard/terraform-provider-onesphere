---
layout: "onesphere"
page_title: "onesphere: memberships"
sidebar_current: "docs-onesphere-memberships"
description: |-
  Adds membership to a project.
---

# onesphere\_membership\

Adds membership to a project.

## Example Usage

```js
# add membership to project
resource "onesphere_membership" "default" {
  membershipname     = "membership_name"
  username           = "user_name"
  membershiprole     = "project-member|project-owner"
  projectname        = "project_name"
}
```

## Argument Reference

The following arguments are supported: 

* `membershipname` - (Required) A unique membership name for the resource.

* `username` - (Required) A unique user name for the resource.

* `membershiprole` - (Required) A membership role project-member|project-owner for the resource.

* `projectname` - (Required) A unique project name for the resource.

---