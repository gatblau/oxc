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

type ModelList struct {
	Values []Model
}

func (list *ModelList) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(list)
	return bytes.NewReader(jsonBytes), err
}

// the Model resource
type Model struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Partition   string `json:"partition"`
	Managed     bool   `json:"managed"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}

// get the Model in the http Response
func (model *Model) decode(response *http.Response) (*Model, error) {
	result := new(Model)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item resource
func (model *Model) uri(baseUrl string) (string, error) {
	if len(model.Key) == 0 {
		return "", fmt.Errorf("the model does not have a key: cannot construct Item resource URI")
	}
	return fmt.Sprintf("%s/model/%s", baseUrl, model.Key), nil
}

// get a JSON bytes reader for the item
func (model *Model) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(model)
	return bytes.NewReader(jsonBytes), err
}

func (model *Model) valid() error {
	if len(model.Key) == 0 {
		return fmt.Errorf("model key is missing")
	}
	if len(model.Name) == 0 {
		return fmt.Errorf("model name is missing")
	}
	return nil
}
