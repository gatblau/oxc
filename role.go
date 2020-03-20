/*
   Onix Config Manager - Terraform Provider
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

package oxc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RoleList struct {
	Values []Role
}

func (list *RoleList) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(list)
	return bytes.NewReader(jsonBytes), err
}

// the Role resource
type Role struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Level       int    `json:"level"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ChangedBy   string `json:"changed_by"`
}

// get the Role in the http Response
func (role *Role) decode(response *http.Response) (*Role, error) {
	result := new(Role)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item resource
func (role *Role) uri(baseUrl string) (string, error) {
	if len(role.Key) == 0 {
		return "", fmt.Errorf("the role does not have a key: cannot construct Role resource URI")
	}
	return fmt.Sprintf("%s/role/%s", baseUrl, role.Key), nil
}

// get a JSON bytes reader for the item
func (role *Role) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(role)
	return bytes.NewReader(jsonBytes), err
}

func (role *Role) valid() error {
	if len(role.Key) == 0 {
		return fmt.Errorf("role key is missing")
	}
	if len(role.Name) == 0 {
		return fmt.Errorf("role name is missing")
	}
	return nil
}
