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

	onesphere "github.com/HewlettPackard/hpe-onesphere-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Exists: resourceUserExists,
		Importer: &schema.ResourceImporter{
			State: resourceUserImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)

	if _, err := config.osClient.GetUserByID(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osUsr := onesphere.UserRequest{
		Name:     d.Get("name").(string),
		Email:    d.Get("email").(string),
		Password: d.Get("password").(string),
		Role:     d.Get("role").(string),
	}

	user, osUsrError := config.osClient.CreateUser(osUsr)
	if osUsrError != nil {
		d.SetId("")
		return osUsrError
	}
	if user.ID != "" {
		d.SetId(user.ID)
	}
	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osUsr, err := config.osClient.GetUserByID(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	if osUsr.ID != "" {
		d.SetId(osUsr.ID)
	}
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osUsr := onesphere.UserRequest{
		Name:     d.Get("name").(string),
		Email:    d.Get("email").(string),
		Password: d.Get("password").(string),
		Role:     d.Get("role").(string),
	}
	user, err := config.osClient.UpdateUser(d.Id(), osUsr)
	if err != nil {
		d.SetId("")
		return nil
	}
	if user.ID != "" {
		d.SetId(user.ID)
	}
	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	err := config.osClient.DeleteUser(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func resourceUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceUserRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
