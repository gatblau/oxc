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

// issue a put http request with the link type data as payload to the resource URI
func (c *Client) PutLinkType(linkType *LinkType) (*Result, error) {
	// validates linkType
	if err := linkType.valid(); err != nil {
		return nil, err
	}

	uri, err := linkType.uri(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// make an http put request to the service
	return c.put(uri, linkType)
}

// issue a delete http request to the resource URI
func (c *Client) DeleteLinkType(linkType *LinkType) (*Result, error) {
	uri, err := linkType.uri(c.BaseURL)

	if err != nil {
		return nil, err
	}

	// make an http delete request to the service
	return c.delete(uri)
}

// issue a get http request to the resource URI
func (c *Client) GetLinkType(linkType *LinkType) (*LinkType, error) {
	uri, err := linkType.uri(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// make an http put request to the service
	result, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	return linkType.decode(result)
}
