// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package onesphereterraform

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	ovMutexKV          = mutexkv.NewMutexKV()
	serverHardwareURIs = make(map[string]bool)
)

// HPE OneSphere credentials for Terraform Provider to connect
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"os_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONESPHERE_OS_USER", ""),
			},
			"os_password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONESPHERE_OS_PASSWORD", nil),
			},
			"os_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONESPHERE_OS_ENDPOINT", nil),
			},
			"os_sslverify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONESPHERE_OS_SSLVERIFY", true),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"onesphere_user":                  resourceUser(),
			"onesphere_project":               resourceProject(),
			"onesphere_service":               resourceService(),
			"onesphere_zone":                  resourceZone(),
			"onesphere_virtualmachineprofile": resourceVirtualMachineProfile(),
			"onesphere_network":               resourceNetwork(),
			"onesphere_deployment":            resourceDeployment(),
			"onesphere_membership":            resourceMembership(),
			"onesphere_membershiprole":        resourceMembershiprole(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		OSUsername:  d.Get("os_username").(string),
		OSPassword:  d.Get("os_password").(string),
		OSEndpoint:  d.Get("os_endpoint").(string),
		OSSSLVerify: d.Get("os_sslverify").(bool),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
