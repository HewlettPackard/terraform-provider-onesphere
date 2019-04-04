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
	"log"
	"os"
	"strings"

	onesphere "github.com/HewlettPackard/hpe-onesphere-go"
	"github.com/HewlettPackard/hpe-onesphere-go/tree/feature/integrate-structs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,
		Exists: resourceDeploymentExists,
		Importer: &schema.ResourceImporter{
			State: resourceDeploymentImport,
		},

		Schema: map[string]*schema.Schema{
			"assignexternalip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publickey": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zonename": {
				Type:     schema.TypeString,
				Required: true,
			},
			"regionname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"projectname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtualmachineprofileid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"networkname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"userdata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceinput": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDeploymentExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)
	
	if _, err := config.osClient.GetDeploymentByID(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func resourceDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	var (
		service   onesphere.Service
		zone      onesphere.Zone
		region    onesphere.Region
		project   onesphere.Project
		network   onesphere.Network
		vmProfile onesphere.VirtualMachineProfile
		err       error
	)
	osNetworkArray := []onesphere.DeploymentNetworks{}
	
	if d.Get("servicename").(string) != "" {
		service, err = config.osClient.GetServiceByName(d.Get("servicename").(string))
		if err != nil {
			//d.SetId("")
			return err
		}
	}

	if d.Get("zonename").(string) != "" {
		zone, err = config.osClient.GetZoneByName(d.Get("zonename").(string))
		if err != nil {
			//d.SetId("")
			return err
		}
		if d.Get("networkname").(string) != "" {
			network, err = config.osClient.GetNetworkByNameAndZoneURI(d.Get("networkname").(string), (zone.URI))
			if err != nil {
				//d.SetId("")
				return err
			}
			n := onesphere.DeploymentNetworks{NetworkURI: network.URI}
			osNetworkArray = append(osNetworkArray, n)
		}
	}

	if d.Get("regionname").(string) != "" {
		region, err = config.osClient.GetRegionByName(d.Get("regionname").(string))
		log.Println("deployment region", region)
		if err != nil {
			//d.SetId("")
			return err
		}
	}

	if d.Get("projectname").(string) != "" {
		project, err = config.osClient.GetProjectByName(d.Get("projectname").(string))
		log.Println("deployment project", project)
		if err != nil {
			//d.SetId("")
			return err
		}
	}

	if d.Get("virtualmachineprofileid").(string) != "" {
		vmProfile, err = config.osClient.GetVirtualMachineProfileByID(d.Get("virtualmachineprofileid").(string))
		log.Println("vmProfile", vmProfile)
		if err != nil {
			d.SetId("")
			return err
		}
	}

	osDeploymentReq := onesphere.DeploymentRequest{
		AssignExternalIP:         d.Get("assignexternalip").(string),
		PublicKey:                d.Get("publickey").(string),
		Name:                     d.Get("name").(string),
		ZoneURI:                  zone.URI,
		ProjectURI:               utils.NewNstring(project.URI),
		RegionURI:                utils.NewNstring(region.URI),
		ServiceURI:               utils.NewNstring(service.URI),
		Networks:                 osNetworkArray,
		VirtualMachineProfileURI: utils.NewNstring(vmProfile.URI),
		Version:                  d.Get("version").(string),
		ServiceInput:             d.Get("serviceinput").(string),
		//AssignExternalIP:         d.Get("assignexternalip").(string),
		//PublicKey:                d.Get("publickey").(string),
		/*Name:       "Pramod-Terraform-Deployment2",
		ZoneURI:    "/rest/zones/75a9ecb1-e39f-414d-b92b-de142c10dde7",
		ProjectURI: "/rest/projects/67698bffdcea42bb8c1979ffd71f076c",
		RegionURI:  utils.NewNstring("/rest/regions/a1de7b65-70f9-4878-a95b-207890e95755"),
		ServiceURI: utils.NewNstring("/rest/services/770a5256-5b05-444b-98aa-ec4c99c4b18c"),
		Networks: []onesphere.DeploymentNetworks{
			{
				NetworkURI: "/rest/networks/0b8f844d-2734-451d-a04a-9eb871808dd0",
			},
		},
		VirtualMachineProfileURI: "/rest/virtual-machine-profiles/2",
		Version:                  d.Get("version").(string),
		ServiceInput:             d.Get("serviceinput").(string),*/
	}
	log.Println("deployment request", osDeploymentReq)
	osDeploymentCreate, osDeploymentError := config.osClient.CreateDeployment(osDeploymentReq)
	log.Println("After osDeployment", osDeploymentCreate.ID) //this should print but it's not
	if osDeploymentError != nil {
		d.SetId("")
		return osDeploymentError
	}
	if osDeploymentCreate.ID != "" {
		d.SetId(osDeploymentCreate.ID)
		//return osDeploymentError
	}
	return nil
}

func resourceDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	log.Println("inside resourceDeploymentRead")
	osDeploymmentRead, err := config.osClient.GetDeploymentByID(d.Id())
	if err != nil {
		//d.SetId("")
		return nil
	}
	if osDeploymmentRead.ID != "" {
	d.SetId(osDeploymmentRead.ID)
	d.Set("name",osDeploymmentRead.Name)
	/*d.Set("networkname",osDeploymmentRead.Name)
	d.Set("projectname",osDeploymmentRead.Name)
	d.Set("regionname",osDeploymmentRead.Name)
	d.Set("servicename",osDeploymmentRead.Name)
	d.Set("zonename",osDeploymmentRead.Name)
	d.Set("virtualmachineprofileid",osDeploymmentRead.memorySizeGB)*/
}
	log.Println("deployment resourceDeploymentRead name", osDeploymmentRead.Name)
	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	log.Println("deployment resourceDeploymentUpdate name")
	
	//updates := []*onesphere.PatchOp{}
	updates := []*onesphere.PatchOp{
		{
		Op: "replace",
		Path: "/name",
		Value: d.Get("name").(string),
	},
	}
	if _, err := config.osClient.UpdateDeployment(d.Id(),updates); err != nil {
		
	}
	return resourceDeploymentRead(d, meta)
}

func resourceDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	log.Println("inside resourceDeploymentDelete")
	err1 := config.osClient.DeleteDeployment(d.Id())
	if err1 != nil {
		return err1
	}
	return nil
}

func resourceDeploymentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)
	log.Println("inside resourceDeploymentImport")
	if err := resourceDeploymentRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
