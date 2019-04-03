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
	"fmt"
	"testing"

	onesphere "github.com/HewlettPackard/hpe-onesphere-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccProject_1(t *testing.T) {
	var project ov.Project

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProject,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProjectExists(
						"onesphere_project.test", &project),
					resource.TestCheckResourceAttr(
						"onesphere_project.test", "name", "Terraform project 1",
					),
				),
			},
			{
				ResourceName:      testAccProject,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckProjectExists(n string, project *ov.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found :%v", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config, err := testProviderConfig()
		if err != nil {
			return err
		}

		testProject, err := config.ovClient.GetProjectByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testProject.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*project = testProject
		return nil
	}
}

func testAccCheckProjectDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "onesphere_project" {
			continue
		}

		testNet, _ := config.ovClient.GetProjectByName(rs.Primary.ID)

		if testNet.Name != "" {
			return fmt.Errorf("Project still exists")
		}
	}

	return nil
}

var testAccProject = `
  resource "onesphere_project" "test" {
	name = "Terraform Project 1"
	description = "terraform project"
    taguris     = ["/rest/tags/environment=demonstration", "/rest/tags/line-of-business=incubation", "/rest/tags/tier=gold"]
  }`
