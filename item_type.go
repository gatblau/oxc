package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// the Item Type resource
type ItemType struct {
	Id           string                 `json:"id"`
	Key          string                 `json:"key"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Filter       map[string]interface{} `json:"filter"`
	MetaSchema   map[string]interface{} `json:"metaSchema"`
	Model        string                 `json:"modelKey"`
	NotifyChange bool                   `json:"notifyChange"`
	Tag          []interface{}          `json:"tag"`
	EncryptMeta  bool                   `json:"encryptMeta"`
	EncryptTxt   bool                   `json:"encryptTxt"`
	Managed      bool                   `json:"managed"`
	Version      int64                  `json:"version"`
	Created      string                 `json:"created"`
	Updated      string                 `json:"updated"`
}

// get the Item Type in the http Response
func (itemType *ItemType) decode(response *http.Response) (*ItemType, error) {
	result := new(ItemType)
	err := json.NewDecoder(response.Body).Decode(result)
	return result, err
}

// get the FQN for the item type resource
func (itemType *ItemType) uri(baseUrl string) (string, error) {
	if len(itemType.Key) == 0 {
		return "", fmt.Errorf("the item does not have a key: cannot construct itemtype resource URI")
	}
	return fmt.Sprintf("%s/itemtype/%s", baseUrl, itemType.Key), nil
}

// get a JSON bytes reader for the itemtype
func (itemType *ItemType) json() (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(itemType)
	return bytes.NewReader(jsonBytes), err
}

func (itemType *ItemType) valid() error {
	if len(itemType.Key) == 0 {
		return fmt.Errorf("item type key is missing")
	}
	if len(itemType.Name) == 0 {
		return fmt.Errorf("item type name is missing")
	}
	return nil
}
