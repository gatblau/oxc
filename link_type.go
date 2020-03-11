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

type LinkType struct {
	Id          string                 `json:"id"`
	Key         string                 `json:"key"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	MetaSchema  map[string]interface{} `json:"metaSchema"`
	Model       string                 `json:"modelKey"`
	Tag         []interface{}          `json:"tag"`
	EncryptMeta bool                   `json:"encryptMeta"`
	EncryptTxt  bool                   `json:"encryptTxt"`
	Managed     bool                   `json:"managed"`
	Version     int64                  `json:"version"`
	Created     string                 `json:"created"`
	Updated     string                 `json:"updated"`
}

// get the Link Type in the http Response
func (linkType *LinkType) decode(response *http.Response) (*LinkType, error) {
	result := new(LinkType)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the link type resource
func (linkType *LinkType) uri(baseUrl string) (string, error) {
	if len(linkType.Key) == 0 {
		return "", fmt.Errorf("the link type does not have a key: cannot construct linktype resource URI")
	}
	return fmt.Sprintf("%s/linktype/%s", baseUrl, linkType.Key), nil
}

// get a JSON bytes reader for the linktype
func (linkType *LinkType) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(linkType)
	return bytes.NewReader(jsonBytes), err
}

func (linkType *LinkType) valid() error {
	if len(linkType.Key) == 0 {
		return fmt.Errorf("link key is missing")
	}
	if len(linkType.Model) == 0 {
		return fmt.Errorf("model is missing")
	}
	return nil
}
