# terraform-provider-onesphere

A Terraform provider for HPE OneSphere

## Installing `terraform-provider-onesphere` with Go

* Install Go 1.11. For previous versions, you may have to set your `$GOPATH` manually, if you haven't done it yet.
* Install Terraform 0.9.x or above [from here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists)
* Download the code by issuing a `go get` command.

```bash
# Download the source code for terraform-provider-onesphere
# and build the needed binary, by saving it inside $GOPATH/bin
$ go get -u github.com/HewlettPackard/terraform-provider-onesphere

# Copy the binary to have it along the terraform binary
$ mv $GOPATH/bin/terraform-provider-onesphere /usr/local/bin/terraform
```

## Example terraform file to provision a virtual machine

```js
# OneSphere Credentials
provider "onesphere" {
  os_username = ONESPHERE_USERNAME
  os_password = ONESPHERE_PASSWORD
  os_endpoint = ONESPHERE_PORTAL_URL
  os_sslverify = true
}

# Create a new OneSphere User
resource "onesphere_user" "default" {
  name     = "user_name"
  email    = "user_email"
  password = "user_password"
  role     = "administrator|analyst|consumer|project-creator"
}

# Create a new OneSphere Project
resource "onesphere_project" "default" {
  name        = "project_name"
  description = "project_description"
  taguris     = ["tag1","tag2","tag3"]
}

# Add membership to Project
resource "onesphere_membership" "default" {
  membershipname     = "membership_name"
  username           = "user_name"
  membershiprole     = "project-member|project-owner"
  projectname        = "project_name"
}

# Add Network to Project
resource "onesphere_network" "default" {
  networkname       = "network_name"
  zonename          = "zone_name"
  operation         = "add|remove"
  projectname       = "project_name"
}

# Create a new OneSphere Deployment
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

More information about how to configure the provider can be found [here](docs/index.html.markdown)

## Resources

Any resource that onesphere can manage is on the roadmap for Terraform to also manage. Below is the current list of resources that Terraform can manage. Open an issue if there is a resource that needs to be developed as soon as possible.

#### [User](docs/r/user.html.markdown)

```js
resource "onesphere_user" "default" {
  name     = "user_name"
  email    = "user_email"
  password = "user_password"
  role     = "administrator|analyst|consumer|project-creator"
}
```

#### [Project](docs/r/project.html.markdown)

```js
resource "onesphere_project" "default" {
  name        = "project_name"
  description = "project_description"
  taguris     = ["tag1","tag2","tag3"]
}
```

#### [Membership](docs/r/membership.html.markdown)

```js
resource "onesphere_membership" "default" {
  membershipname     = "membership_name"
  username           = "user_name"
  membershiprole     = "project-member|project-owner"
  projectname        = "project_name"
}
```

#### [Network add|remove](docs/r/network.html.markdown)

```js
resource "onesphere_network" "default" {
  networkname       = "network_name"
  zonename          = "zone_name"
  operation         = "add|remove"
  projectname       = "project_name"
}
```

#### [Deployment](docs/r/deployment.markdown)

```js
resource "onesphere_deployment" "default" {
  name                          = "deployment_name"
  zonename                      = "synergy-vsan-deic"
  regionname                    = "emea-mougins-fr"
  servicename                   = "rhel61forcsa"
  projectname                   = "terraform"
  virtualmachineprofileid       = "1|2|3|4"
  networkname                   = "network_name"
  serviceinput                  = "--port=80 --target-port=8080 --type=LoadBalancer"
}
```

### License

This project is licensed under the Apache 2.0 license.
