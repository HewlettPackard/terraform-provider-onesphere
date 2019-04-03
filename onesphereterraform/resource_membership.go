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

func resourceMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceMembershipCreate,
		Read:   resourceMembershipRead,
		Update: resourceMembershipUpdate,
		Delete: resourceMembershipDelete,
		Exists: resourceMembershipExists,
		Importer: &schema.ResourceImporter{
			State: resourceMembershipImport,
		},

		Schema: map[string]*schema.Schema{
			"membershipname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"membershiprole": {
				Type:     schema.TypeString,
				Required: true,
			},
			"projectname": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMembershipExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)

	if _, err := config.osClient.GetMembershipByID(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osUsr, usrerr := config.osClient.GetUserByName(d.Get("username").(string))
	if usrerr != nil {
		d.SetId("")
		return nil
	}
	osProj, projerr := config.osClient.GetProjectByName(d.Get("projectname").(string))
	if projerr != nil {
		d.SetId("")
		return nil
	}
	osMembershiprole, memerr := config.osClient.GetMembershipRoleByName(d.Get("membershiprole").(string))
	if memerr != nil {
		d.SetId("")
		return nil
	}

	osMemb := onesphere.Membership{
		UserURI:           osUsr.URI,
		MembershipRoleURI: osMembershiprole.URI,
		ProjectURI:        osProj.URI,
	}

	Mem, osMemError := config.osClient.CreateMembership(osMemb)

	if osMemError != nil {
		d.SetId("1234589654")
		d.SetId("")
		return osMemError
	}
	if Mem.ID != "" {
		//d.SetId(Mem.ID)
	}
	//return resourceMembershipRead(d, meta)
	return nil
}

func resourceMembershipRead(d *schema.ResourceData, meta interface{}) error {
	/*config := meta.(*Config)

	osProj, err := config.osClient.GetMembershipByName(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))*/

	return nil
}

func resourceMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}

func resourceMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	osUsr, usrerr := config.osClient.GetUserByName(d.Get("username").(string))
	osProj, projerr := config.osClient.GetProjectByName(d.Get("projectname").(string))
	osMembershiprole, memerr := config.osClient.GetMembershipRoleByName(d.Get("membershiprole").(string))
	if usrerr != nil {
		d.SetId("")
		return nil
	}
	if projerr != nil {
		d.SetId("")
		return nil
	}
	if memerr != nil {
		d.SetId("")
		return nil
	}
	osMemb := onesphere.Membership{
		UserURI:           osUsr.URI,
		MembershipRoleURI: osMembershiprole.URI,
		ProjectURI:        osProj.URI,
	}

	osMemError := config.osClient.DeleteMembershipByID(osMemb.ProjectURI)
	if osMemError != nil {
		return osMemError
	}
	return nil
}

func resourceMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceMembershipRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
