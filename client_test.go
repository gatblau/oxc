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
	"fmt"
	"testing"
)

var client Client

func init() {
	client = Client{BaseURL: "http://localhost:8080"}

	// test using a basic authentication token
	t := client.newBasicToken("admin", "0n1x")

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

	client.setAuthToken(t)
}

func checkResult(result *Result, err error, msg string, t *testing.T) {
	if err != nil {
		t.Error(msg)
	}
	if result.Error {
		t.Error(fmt.Sprintf("%s: %s", msg, result.Message))
	}
}

func checkError(err error, msg string, t *testing.T) {
	if err != nil {
		t.Error(msg, err)
	}
}

func TestOnixClient_Put(t *testing.T) {
	msg := "create test_model failed"
	model := &Model{
		Key:         "test_model",
		Name:        "Test Model",
		Description: "Test Model",
	}
	result, err := client.putModel(model)
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
	result, err = client.putItemType(itemType)
	checkResult(result, err, "create test_item_type failed", t)

	item_1 := &Item{
		Key:         "item_1",
		Name:        "Item 1",
		Description: "Test Item 1",
		Status:      1,
		Type:        "test_item_type",
		Txt:         "This is a test text configuration.",
	}
	result, err = client.putItem(item_1)
	checkResult(result, err, "create item_1 failed", t)

	item_2 := &Item{
		Key:         "item_2",
		Name:        "Item 2",
		Description: "Test Item 2",
		Status:      2,
		Type:        "test_item_type",
	}
	result, err = client.putItem(item_2)
	checkResult(result, err, "create item_2 failed", t)

	link_type := &LinkType{
		Key:         "test_link_type",
		Name:        "Test Link Type",
		Description: "Test Link Type",
		Model:       "test_model",
	}
	result, err = client.putLinkType(link_type)
	checkResult(result, err, "create test_link_type failed", t)

	link_rule := &LinkRule{
		Key:              "test_link_rule_1",
		Name:             "Test Item Type to Test Item Type rule",
		Description:      "Allow to connect two items of type test_item_type.",
		LinkTypeKey:      "test_link_type",
		StartItemTypeKey: "test_item_type",
		EndItemTypeKey:   "test_item_type",
	}
	result, err = client.putLinkRule(link_rule)
	checkResult(result, err, "create test_item_type->test_item_type rule failed", t)

	link := &Link{
		Key:          "test_link_1",
		Description:  "Test Link 1",
		Type:         "test_link_type",
		StartItemKey: "item_1",
		EndItemKey:   "item_2",
	}
	result, err = client.putLink(link)
	checkResult(result, err, "create link_1 failed", t)
}
