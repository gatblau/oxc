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
package oxc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Link struct {
	Id           string                 `json:"id"`
	Key          string                 `json:"key"`
	Description  string                 `json:"description"`
	Type         string                 `json:"type"`
	Tag          []interface{}          `json:"tag"`
	Meta         map[string]interface{} `json:"meta"`
	Attribute    map[string]interface{} `json:"attribute"`
	StartItemKey string                 `json:"startItemKey"`
	EndItemKey   string                 `json:"endItemKey"`
	Version      int64                  `json:"version"`
	Created      string                 `json:"created"`
	Updated      string                 `json:"updated"`
}

// get the Link in the http Response
func (link *Link) decode(response *http.Response) (*Link, error) {
	result := new(Link)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the link type resource
func (link *Link) uri(baseUrl string) (string, error) {
	if len(link.Key) == 0 {
		return "", fmt.Errorf("the link does not have a key: cannot construct link resource URI")
	}
	return fmt.Sprintf("%s/link/%s", baseUrl, link.Key), nil
}

// get a JSON bytes reader for the link
func (link *Link) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(link)
	return bytes.NewReader(jsonBytes), err
}

func (link *Link) valid() error {
	if len(link.Key) == 0 {
		return fmt.Errorf("link key is missing")
	}
	if len(link.Type) == 0 {
		return fmt.Errorf("link type is missing")
	}
	if len(link.StartItemKey) == 0 {
		return fmt.Errorf("start item key is missing")
	}
	if len(link.EndItemKey) == 0 {
		return fmt.Errorf("end item key is missing")
	}
	return nil
}
