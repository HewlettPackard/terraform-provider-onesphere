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
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVirtualMachineProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMProfileCreate,
		Read:   resourceVMProfileRead,
		Update: resourceVMProfileUpdate,
		Delete: resourceVMProfileDelete,
		Exists: resourceVMProfileExists,
		Importer: &schema.ResourceImporter{
			State: resourceVMProfileImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVMProfileExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)

	if _, err := config.osClient.GetVirtualMachineProfileByID(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceVMProfileCreate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return resourceDeploymentRead(d, meta)
}

func resourceVMProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osVMProfileList, err := config.osClient.GetVirtualMachineProfiles(d.Get("name").(string))
	if err != nil {
		//d.SetId("")
		return nil
	}
	d.Set("name", osVMProfileList.Name)
	return nil
}

func resourceVMProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}

func resourceVMProfileDelete(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}

func resourceVMProfileImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceVMProfileRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
