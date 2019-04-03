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
	"errors"

	onesphere "github.com/HewlettPackard/hpe-onesphere-go"
)

//Config structure
type Config struct {
	OSUsername  string
	OSPassword  string
	OSEndpoint  string
	OSSSLVerify bool

	osClient *onesphere.Client
}

var ErrConfigNotInitialized = errors.New("config not initialized!")

func (c *Config) loadAndValidate() error {
	if c == nil {
		return ErrConfigNotInitialized
	}

	client, err := onesphere.Connect(c.OSEndpoint, c.OSUsername, c.OSPassword)

	c.osClient = client

	//session, err := c.osClient.SessionLogin()
	if err != nil {
		return err
	}

	//	c.osClient.APIKey = session.Token

	return nil
}
