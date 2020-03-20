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

type PartitionList struct {
	Values []Partition
}

func (list *PartitionList) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(list)
	return bytes.NewReader(jsonBytes), err
}

// the Partition resource
type Partition struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Version     int64  `json:"version"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ChangedBy   string `json:"changed_by"`
}

// get the Partition in the http Response
func (partition *Partition) decode(response *http.Response) (*Partition, error) {
	result := new(Partition)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item resource
func (partition *Partition) uri(baseUrl string) (string, error) {
	if len(partition.Key) == 0 {
		return "", fmt.Errorf("the partition does not have a key: cannot construct Partition resource URI")
	}
	return fmt.Sprintf("%s/partition/%s", baseUrl, partition.Key), nil
}

// get a JSON bytes reader for the item
func (partition *Partition) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(partition)
	return bytes.NewReader(jsonBytes), err
}

func (partition *Partition) valid() error {
	if len(partition.Key) == 0 {
		return fmt.Errorf("partition key is missing")
	}
	if len(partition.Name) == 0 {
		return fmt.Errorf("partition name is missing")
	}
	return nil
}
