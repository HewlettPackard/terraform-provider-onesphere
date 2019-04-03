---
layout: "onesphere"
page_title: "Provider: OneSphere"
sidebar_current: "docs-onesphere-index"
description: |-
  The OneSphere provider is used to interact with your OneSphere portal. The provider needs to be configured with the proper credentials before it can be used. 
---

# OneSphere Provider

 The OneSphere provider is used to interact with [HPE OneSphere](https://www.hpe.com/us/en/solutions/cloud/hybrid-it-management.html).
 The provider needs to be configured with the proper credentials before it can be used.

## Example Usage

```js
//Configure the OneSphere Provider
provider "onesphere" {
  os_username = ONESPHERE_USERNAME
  os_password = ONESPHERE_PASSWORD
  os_endpoint = ONESPHERE_PORTAL_URL
  os_sslverify = true
}

//Create a new OneSphere user
resource "onesphere_user" "default" {
  name     = "user_name"
  email    = "user_email"
  password = "user_password"
  role     = "administrator|analyst|consumer|project-creator"
}
```

## Authentication

The OneSphere provider supports static credentials and environment variables.

## Configuration Reference

The following keys can be used to configure the provider.

* `os_username` - (Optional) This is the OneSphere username. 
  It must be provided or sourced from OneSphere_OS_USER environment variable.

* `os_password` - (Optional) This is the OneSphere password. 
  It must be provided or sourced from OneSphere_OS_PASSWORD environment variable.
  
* `os_endpoint` - (Optional) This is the OneSphere URL.
  It must be provided or sourced from OneSphere_OS_ENDPOINT environment variable.

* `os_sslverify` - (Optional) This is a boolean value for whether ssl is enabled. 
  It must be provided or sourced from OneSphere_OS_SSLVERIFY environment variable.
