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

type UserList struct {
	Values []User
}

func (list *UserList) json() (*bytes.Reader, error) {
	jsonBytes, err := jsonBytes(list)
	return bytes.NewReader(jsonBytes), err
}

// the Role resource
type User struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Pwd       string `json:"pwd"`
	Expires   string `json:"expires"`
	Version   int64  `json:"version"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	ChangedBy string `json:"changedBy"`
}

// Get the Role in the http Response
func (user *User) decode(response *http.Response) (*User, error) {
	result := new(User)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// Get the FQN for the item resource
func (user *User) uri(baseUrl string) (string, error) {
	if len(user.Key) == 0 {
		return "", fmt.Errorf("the user does not have a key: cannot construct User resource URI")
	}
	return fmt.Sprintf("%s/user/%s", baseUrl, user.Key), nil
}

// Get a JSON bytes reader for the entity
func (user *User) json() (*bytes.Reader, error) {
	jsonBytes, err := user.bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(*jsonBytes), err
}

// Get a []byte representing the entity
func (user *User) bytes() (*[]byte, error) {
	b, err := jsonBytes(user)
	return &b, err
}

func (user *User) valid() error {
	if len(user.Key) == 0 {
		return fmt.Errorf("user key is missing")
	}
	if len(user.Name) == 0 {
		return fmt.Errorf("user name is missing")
	}
	if len(user.Name) == 0 {
		return fmt.Errorf("user name is missing")
	}
	return nil
}