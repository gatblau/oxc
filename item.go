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

type ItemList struct {
	Values []Item
}

func (list *ItemList) json() (*bytes.Reader, error) {
	jsonBytes, err := jsonBytes(list)
	return bytes.NewReader(jsonBytes), err
}

// the Item resource
type Item struct {
	Key         string                 `json:"key"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Status      int                    `json:"status"`
	Type        string                 `json:"type"`
	Tag         []interface{}          `json:"tag"`
	Meta        map[string]interface{} `json:"meta"`
	Txt         string                 `json:"txt"`
	Attribute   map[string]interface{} `json:"attribute"`
	Partition   string                 `json:"partition"`
	Version     int64                  `json:"version"`
	Created     string                 `json:"created"`
	Updated     string                 `json:"updated"`
	EncKeyIx    int64                  `json:"encKeyIx"`
	ChangedBy   string                 `json:"changedBy"`
}

// get the Item in the http Response
func (item *Item) decode(response *http.Response) (*Item, error) {
	result := new(Item)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the ItemList in the http Response
func (item *Item) decodeList(response *http.Response) (*ItemList, error) {
	result := new(ItemList)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item resource
func (item *Item) uri(baseUrl string) (string, error) {
	if len(item.Key) == 0 {
		return "", fmt.Errorf("the item does not have a key: cannot construct Item resource URI")
	}
	return fmt.Sprintf("%s/item/%s", baseUrl, item.Key), nil
}

// get a JSON bytes reader for the entity
func (item *Item) json() (*bytes.Reader, error) {
	jsonBytes, err := item.bytes()
	return bytes.NewReader(*jsonBytes), err
}

// get a []byte representing the entity
func (item *Item) bytes() (*[]byte, error) {
	bytes, err := jsonBytes(item)
	return &bytes, err
}

func (item *Item) valid() error {
	if len(item.Key) == 0 {
		return fmt.Errorf("item key is missing")
	}
	if len(item.Name) == 0 {
		return fmt.Errorf("item name is missing")
	}
	return nil
}

// get the FQN for the item / children resource
func (item *Item) uriItemChildren(baseUrl string) (string, error) {
	if len(item.Key) == 0 {
		return "", fmt.Errorf("the item does not have a key: cannot construct Item/Children resource URI")
	}
	return fmt.Sprintf("%s/item/%s/children", baseUrl, item.Key), nil
}
