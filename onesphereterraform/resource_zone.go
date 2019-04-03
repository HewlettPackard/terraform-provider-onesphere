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

func resourceZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneCreate,
		Read:   resourceZoneRead,
		Update: resourceZoneUpdate,
		Delete: resourceZoneDelete,
		Exists: resourceZoneExists,
		Importer: &schema.ResourceImporter{
			State: resourceZoneImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZoneExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)

	if _, err := config.osClient.GetZoneByID(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceZoneCreate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	return resourceZoneRead(d, meta)
}

func resourceZoneRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	osZone, err := config.osClient.GetZoneByName(d.Get("zonename").(string))
	if err != nil {
		d.SetId(osZone.ID)
		return nil
	}
	return nil
}

func resourceZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}

func resourceZoneDelete(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}

func resourceZoneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceZoneRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
