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
	"fmt"
	"testing"
)

var client Client

func init() {
	client = Client{BaseURL: "http://localhost:8080"}

	// test using a basic authentication token
	t := client.NewBasicToken("admin", "0n1x")

	// uncomment below & reset configuration vars
	// to test using using an OAuth bearer token

	//t, err := client.getBearerToken(
	//	"https://dev-447786.okta.com/oauth2/default/v1/token",
	//	"0oalyh...356",
	//	"Tsed........OP0oEf9H7",
	//	"user@email.com",
	//	"user_pwd_xxxxx")
	//
	//if err != nil {
	//	panic(err)
	//}

	client.SetAuthToken(t)
}

func checkResult(result *Result, err error, msg string, t *testing.T) {
	if err != nil {
		t.Error(msg)
	} else if result != nil {
		if result.Error {
			t.Error(fmt.Sprintf("%s: %s", msg, result.Message))
		} else if result.Operation == "L" {
			t.Error(fmt.Sprintf("Fail to update - Locked entity: %s", result.Ref))
		}
	}
}

func TestOnixClient_Put(t *testing.T) {
	// clear all data!
	result, err := client.Clear()
	checkResult(result, err, "failed to clear database", t)

	msg := "create test_model failed"
	model := &Model{
		Key:         "test_model",
		Name:        "Test Model",
		Description: "Test Model",
	}
	result, err = client.PutModel(model)
	checkResult(result, err, msg, t)

	itemType := &ItemType{
		Key:          "test_item_type",
		Name:         "Test Item Type",
		Description:  "Test Item Type",
		Model:        "test_model",
		EncryptMeta:  false,
		EncryptTxt:   true,
		Managed:      true,
		NotifyChange: true,
	}
	result, err = client.PutItemType(itemType)
	checkResult(result, err, "create test_item_type failed", t)

	itemTypeAttr := &ItemTypeAttribute{
		Key:         "test_item_type_attr_1",
		Name:        "CPU",
		Description: "Description for test_item_type_attr_1",
		Type:        "integer",
		DefValue:    "2",
		Managed:     false,
		Required:    false,
		Regex:       "",
		ItemTypeKey: "test_item_type",
	}

	result, err = client.PutItemTypeAttr(itemTypeAttr)
	checkResult(result, err, "create test_item_type_attr_1 failed", t)

	item_1 := &Item{
		Key:         "item_1",
		Name:        "Item 1",
		Description: "Test Item 1",
		Status:      1,
		Type:        "test_item_type",
		Txt:         "This is a test text configuration.",
		Attribute:   map[string]interface{}{"CPU": 5},
	}
	result, err = client.PutItem(item_1)
	checkResult(result, err, "create item_1 failed", t)

	item_2 := &Item{
		Key:         "item_2",
		Name:        "Item 2",
		Description: "Test Item 2",
		Status:      2,
		Type:        "test_item_type",
		Attribute:   map[string]interface{}{"CPU": 2},
	}
	result, err = client.PutItem(item_2)
	checkResult(result, err, "create item_2 failed", t)

	link_type := &LinkType{
		Key:         "test_link_type",
		Name:        "Test Link Type",
		Description: "Test Link Type",
		Model:       "test_model",
	}
	result, err = client.PutLinkType(link_type)
	checkResult(result, err, "create test_link_type failed", t)

	link_rule := &LinkRule{
		Key:              "test_link_rule_1",
		Name:             "Test Item Type to Test Item Type rule",
		Description:      "Allow to connect two items of type test_item_type.",
		LinkTypeKey:      "test_link_type",
		StartItemTypeKey: "test_item_type",
		EndItemTypeKey:   "test_item_type",
	}
	result, err = client.PutLinkRule(link_rule)
	checkResult(result, err, "create test_item_type->test_item_type rule failed", t)

	link := &Link{
		Key:          "test_link_1",
		Description:  "Test Link 1",
		Type:         "test_link_type",
		StartItemKey: "item_1",
		EndItemKey:   "item_2",
	}
	result, err = client.PutLink(link)
	checkResult(result, err, "create link_1 failed", t)

	data := getData()
	result, err = client.PutData(data)
	if err != nil {
		t.Error(err)
	}
	if result.Error {
		t.Error(result.Message)
	}

	list, err := client.GetItemChildren(&Item{Key: "item_1"})
	if err != nil {
		t.Error(err)
	}
	if len(list.Values) == 0 {
		t.Error("no value in list")
	}
}

func getData() *GraphData {
	return &GraphData{
		Models: []Model{
			Model{
				Key:         "TERRA",
				Name:        "Terraform Model",
				Description: "Defines the item and link types that describe Terraform resources.",
			},
		},
		ItemTypes: []ItemType{
			ItemType{
				Key:         "TF_STATE",
				Name:        "Terraform State",
				Description: "State about a group of managed infrastructure and configuration resources. This state is used by Terraform to map real world resources to your configuration, keep track of metadata, and to improve performance for large infrastructures.",
				Model:       "TERRA",
			},
			ItemType{
				Key:         "TF_RESOURCE",
				Name:        "Terraform Resource",
				Description: "Each resource block describes one or more infrastructure objects, such as virtual networks, compute instances, or higher-level components such as DNS records.",
				Model:       "TERRA",
			},
		},
		ItemTypeAttributes: []ItemTypeAttribute{
			ItemTypeAttribute{
				Key:         "TF_ITEM_ATTR_MODE",
				Name:        "mode",
				Description: "Whether the resource is a data source or a managed resource.",
				Type:        "string",
				ItemTypeKey: "TF_RESOURCE",
				Required:    true,
			},
			ItemTypeAttribute{
				Key:         "TF_ITEM_ATTR_TYPE",
				Name:        "type",
				Description: "The resource type.",
				Type:        "string",
				ItemTypeKey: "TF_RESOURCE",
				Required:    true,
			},
			ItemTypeAttribute{
				Key:         "TF_ITEM_ATTR_PROVIDER",
				Name:        "provider",
				Description: "The provider used to manage this resource.",
				Type:        "string",
				ItemTypeKey: "TF_RESOURCE",
				Required:    true,
			},
		},
		LinkTypes: []LinkType{
			LinkType{
				Key:         "TF_STATE_LINK",
				Name:        "Terraform State Link Type",
				Description: "Links Terraform resources that are part of a state.",
				Model:       "TERRA",
			},
		},
		LinkRules: []LinkRule{
			LinkRule{
				Key:              fmt.Sprintf("%s->%s", "TF_STATE", "TF_RESOURCE"),
				Name:             "Terraform State to Resource Rule",
				Description:      "Allow the linking of a Terraform State item to one or more Terraform Resource items using Terraform State Links.",
				LinkTypeKey:      "TF_STATE_LINK",
				StartItemTypeKey: "TF_STATE",
				EndItemTypeKey:   "TF_RESOURCE",
			},
		},
	}
}
