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
	"os"
	"log"

	onesphere "github.com/HewlettPackard/hpe-onesphere-go"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,
		Exists: resourceProjectExists,
		Importer: &schema.ResourceImporter{
			State: resourceProjectImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"taguris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceProjectExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	config := meta.(*Config)
	osProjct, err := config.osClient.GetProjectByID(d.Id(), "full")

	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	d.SetId(osProjct.ID)
	return true, nil
}

func resourceProjectCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	rawTagUris := d.Get("taguris").(*schema.Set).List()
	tagUris := make([]string, len(rawTagUris))
	for i, raw := range rawTagUris {
		tagUris[i] = raw.(string)
	}

	osProj := onesphere.ProjectRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		TagUris:     tagUris,
	}

	project, osProjError := config.osClient.CreateProject(osProj)

	if osProjError != nil {
		d.SetId("")
		return osProjError
	}
	if project.ID != "" {
		d.SetId(project.ID)
	}
	return resourceProjectRead(d, meta)
}

func resourceProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	if d.Get("name") != "" {
		osProj, err := config.osClient.GetProjectByName(d.Get("name").(string))
		if err != nil {
			d.SetId("")
			return nil
		}
		d.SetId(osProj.ID)
	}
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)

	rawTagUris := d.Get("taguris").(*schema.Set).List()
	tagUris := make([]string, len(rawTagUris))
	for i, raw := range rawTagUris {
		tagUris[i] = raw.(string)
	}

	newProj := onesphere.ProjectRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		TagUris: tagUris,
	}
	log.Println("before UpdateProject call")
	project, err := config.osClient.UpdateProject(d.Get("id").(string), newProj)
	log.Println("After UpdateProject call")
	if err != nil {
		return err
	}
	if project.ID != "" {
		d.SetId(project.ID)
	}

	return resourceProjectRead(d, meta)
}

func resourceProjectDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	err1 := config.osClient.DeleteProject(d.Id())
	if err1 != nil {
		return err1
	}
	return nil
}

func resourceProjectImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceProjectRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
