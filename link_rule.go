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

type LinkRule struct {
	Id               string `json:"id"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	LinkTypeKey      string `json:"linkTypeKey"`
	StartItemTypeKey string `json:"startItemTypeKey"`
	EndItemTypeKey   string `json:"endItemTypeKey"`
	Version          int64  `json:"version"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
}

// get the Link Rule in the http Response
func (linkRule *LinkRule) decode(response *http.Response) (*LinkRule, error) {
	result := new(LinkRule)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item type resource
func (rule *LinkRule) uri(baseUrl string) (string, error) {
	if len(rule.Key) == 0 {
		return "", fmt.Errorf("the linkrule does not have a key: cannot construct linkrule resource URI")
	}
	return fmt.Sprintf("%s/linkrule/%s", baseUrl, rule.Key), nil
}

// get a JSON bytes reader for the linkrule
func (rule *LinkRule) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(rule)
	return bytes.NewReader(jsonBytes), err
}

func (rule *LinkRule) valid() error {
	if len(rule.Key) == 0 {
		return fmt.Errorf("item type attribute key is missing")
	}
	if len(rule.StartItemTypeKey) == 0 {
		return fmt.Errorf("start item type key is missing")
	}
	if len(rule.EndItemTypeKey) == 0 {
		return fmt.Errorf("end item type key is missing")
	}
	return nil
}
