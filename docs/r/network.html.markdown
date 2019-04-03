---
layout: "onesphere"
page_title: "onesphere: network"
sidebar_current: "docs-onesphere-network"
description: |-
  Add or Remove network access to a project.
---

# onesphere\_network\

Updates a network.

## Example Usage

```js
resource "onesphere_network" "default" {
  networkname       = "network_name"
  zonename          = "zone_name"
  operation         = "add|remove"
  projectname       = "project_name"
}

```

## Argument Reference

The following arguments are supported: 

* `networkname` - (Required) A unique network name for the resource.

* `zonename` - (Required) A unique zone name for the resource.

* `operation` - (Required) ADD|REMOVE network access for a project.

* `projectname` - (Required) A project name which needs the network access.

---