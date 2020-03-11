/*
   Onix Configuration Manager - Web Api go client
   Copyright (c) 2018-2020 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// the Privilege resource
type Privilege struct {
	Id        string `json:"id"`
	Key       string `json:"key"`
	Role      string `json:"roleKey"`
	Partition string `json:"partitionKey"`
	Create    bool   `json:"canCreate"`
	Read      bool   `json:"canRead"`
	Delete    bool   `json:"canDelete"`
	Version   int64  `json:"version"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

// get the Privilege in the http Response
func (privilege *Privilege) decode(response *http.Response) (*Privilege, error) {
	result := new(Privilege)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the privilege resource
func (privilege *Privilege) uri(baseUrl string) (string, error) {
	if len(privilege.Key) == 0 {
		return "", fmt.Errorf("the privilege does not have a key: cannot construct privilege resource URI")
	}
	return fmt.Sprintf("%s/privilege/%s", baseUrl, privilege.Key), nil
}

// get a JSON bytes reader for the privilege
func (privilege *Privilege) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(privilege)
	return bytes.NewReader(jsonBytes), err
}

func (privilege *Privilege) valid() error {
	if len(privilege.Key) == 0 {
		return fmt.Errorf("privilege key is missing")
	}
	if len(privilege.Role) == 0 {
		return fmt.Errorf("privilege role is missing")
	}
	if len(privilege.Partition) == 0 {
		return fmt.Errorf("privilege partition is missing")
	}
	return nil
}
