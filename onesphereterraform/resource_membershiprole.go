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
	//onesphere "github.com/HewlettPackard/hpe-onesphere-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMembershiprole() *schema.Resource {
	return &schema.Resource{
		Create: resourceMembershiproleCreate,
		Read:   resourceMembershiproleRead,
		Update: resourceMembershiproleUpdate,
		Delete: resourceMembershiproleDelete,
		Exists: resourceMembershiproleExists,
		Importer: &schema.ResourceImporter{
			State: resourceMembershiproleImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"displayname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMembershiproleExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	//config := meta.(*Config)
	return false, nil
}

func resourceMembershiproleCreate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return nil
}

func resourceMembershiproleRead(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return nil
}

func resourceMembershiproleUpdate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return nil
}

func resourceMembershiproleDelete(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return nil
}

func resourceMembershiproleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	//config := meta.(*Config)
	return nil, nil
}
