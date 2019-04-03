---
layout: "onesphere"
page_title: "onesphere: deployment"
sidebar_current: "docs-onesphere-deployment"
description: |-
  Creates a deployment.
---

# onesphere\_deployment\

Creates a deployment.

## Example Usage

```js
resource "onesphere_deployment" "default" {
  name                          = "deployment_name"
  zonename                      = "zone_name"
  regionname                    = "region_name"
  servicename                   = "service_name"
  projectname                   = "project_name"
  virtualmachineprofileid       = "1|2|3|4"
  networkname                   = "network_name"
  serviceinput                  = "--port=80 --target-port=8080 --type=LoadBalancer"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique deployment name for the resource.

* `zonename` - (Required) A unique zone name for the resource.

* `regionname` - (Required) A unique region name for the resource.

* `servicename` - (Required) A unique service name for the resource.

* `projectname` - (Required) A unique project name for the resource.

* `virtualmachineprofileid` - (Required) A unique virtual machine profile id for the resource.

* `networkname` - (Required) A unique network name for the resource.

* `serviceinput` - (Required) A service input field for application port and type fields for the resource.

---