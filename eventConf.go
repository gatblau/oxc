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
	"crypto/tls"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"os"
)

// configuration for the event manager (mqtt broker)
type EventConfig struct {
	// the MQTT server url
	server string
	// the item type for which to get notification changes (itemInstance must be empty)
	itemType string
	// the item instance for which to get notification changes (itemType must be empty)
	itemInstance string
	// the quality of service for message delivery - 0: at most once, 1: at least once, 2: exactly once
	qos int
	// authentication username
	username string
	// authentication password
	password string
	// skip tls certificate verification
	insecureSkipVerify bool
	// the policy the server will follow for TLS Client Authentication
	clientAuthType tls.ClientAuthType
	// a function to process received messages
	msgReceived MQTT.MessageHandler
}

func (c *EventConfig) hasCredentials() bool {
	return len(c.username) > 0 && len(c.password) > 0
}

func (c *EventConfig) topic() string {
	if len(c.itemInstance) > 0 {
		return fmt.Sprintf("II_%s", c.itemInstance)
	}
	return fmt.Sprintf("IT_%s", c.itemType)
}

// unique identifier for the client
func (c *EventConfig) clientId() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}
	return fmt.Sprintf("%s-%s-%s", c.topic(), hostname, uuid.New())
}

// check the configuration is valid
func (c *EventConfig) isValid() (bool, error) {
	if len(c.server) == 0 {
		return false, errors.New("server property not provided")
	}
	if len(c.itemInstance) > 0 && len(c.itemType) > 0 {
		return false, errors.New("itemType and itemInstance both have values, only one is allowed")
	}
	if len(c.itemInstance) == 0 && len(c.itemType) == 0 {
		return false, errors.New("itemType and itemInstance do not have values, one is required")
	}
	if len(c.username) > 0 && len(c.password) == 0 {
		return false, errors.New("username with no password, provide password")
	}
	if c.msgReceived == nil {
		return false, errors.New("a handler for received messages must be provided")
	}
	return true, nil
}
