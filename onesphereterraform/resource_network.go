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

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: resourceNetworkDelete,
		Exists: resourceNetworkExists,
		Importer: &schema.ResourceImporter{
			State: resourceNetworkImport,
		},

		Schema: map[string]*schema.Schema{
			"networkname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"projectname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"regionname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zonename": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNetworkExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)
	
	osNetwk, err := config.osClient.GetNetworkByID(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	d.SetId(osNetwk.ID)
	return true, nil
}

func resourceNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)
	
	return nil
}

func resourceNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	if d.Get("zonename") != "" {
		osZone, err := config.osClient.GetZoneByName(d.Get("zonename").(string))
		if err != nil {
			//d.SetId(osNetwork.ID)
			return nil
		}
		if osZone.URI != "" && d.Get("networkname") != "" {
			osNetwork, neterr := config.osClient.GetNetworkByNameAndZoneURI(d.Get("networkname").(string), (osZone.URI))
			if neterr != nil {
				d.SetId("")
				return nil
			}
			d.SetId(osNetwork.ID)
		}
	}
	return nil
}

func resourceNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	
	if d.Get("projectname") != "" {
		osProj, osProjerr := config.osClient.GetProjectByName(d.Get("projectname").(string))
		if osProjerr != nil {
			d.SetId("")
			return nil
		}
	
		/*if d.Get("regionname") != "" {
			osRegion, osRegionerr := config.osClient.GetRegionByName(d.Get("regionname").(string))
			if osRegionerr != nil {
				d.SetId("")
				return nil
			}
			log.Println("After osRegion resourceNetworkUpdate ", osRegion.Name)
		}*/
		if d.Get("zonename") != "" {
			osZone, osZoneerr := config.osClient.GetZoneByName(d.Get("zonename").(string))
			if osZoneerr != nil {
				d.SetId("")
				return osZoneerr
			}
			if osZone.URI != "" && d.Get("networkname") != "" {
				//osNetwork, osNetworkerr := config.osClient.GetNetworkByNameAndZoneURI(d.Get("networkname").(string), osZone.URI)
				//osNetwork, osNetworkerr := config.osClient.GetNetworkByID(d.Id())
				osNetwork, osNetworkerr := config.osClient.GetNetworkByNameAndZoneURI(d.Get("networkname").(string), osZone.URI)
				if osNetworkerr != nil {
					d.SetId("")
					return osNetworkerr
				}
				//d.SetId(osNetwork.ID)
				//osNetwork, err := config.osClient.GetNetworkByName(d.Get("networkname").(string))
				osNet := []*onesphere.PatchOp{
					{
						Op:    d.Get("operation").(string),
						Path:  "projectUris",
						Value: osProj.ID,
					},
				}
				osntwrk, osNetError := config.osClient.UpdateNetwork(osNetwork.ID, osNet)
				if osNetError != nil {
					return osNetError
				}
				d.SetId(osntwrk.ID)
			}
		}
	}
	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(*Config)

	return nil
}
func resourceNetworkImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceNetworkRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
