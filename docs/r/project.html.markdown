---
layout: "onesphere"
page_title: "onesphere: project"
sidebar_current: "docs-onesphere-project"
description: |-
  Creates a project.
---

# onesphere\_project\

Creates a project.

## Example Usage

```js
resource "onesphere_project" "default" {
  name        = "project_name"
  description = "project_description"
  taguris     = ["tag1","tag2","tag3"]
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique project name for the resource.

* `description` - (Required) Description for the resource.

* `taguris` - (Required) Tag URI's for the resource.

---