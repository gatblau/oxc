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

// issue a Put http request with the Item data as payload to the resource URI
func (c *Client) PutItem(item *Item) (*Result, error) {
	// validates item
	if err := item.valid(); err != nil {
		return nil, err
	}

	uri, err := item.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}

	// make an http Put request to the service
	resp, err := c.Put(uri, item, c.addHttpHeaders)

	return newResult(resp, err)
}

// issue a Delete http request to the resource URI
func (c *Client) DeleteItem(item *Item) (*Result, error) {
	uri, err := item.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}

	// make an http Delete request to the service
	resp, err := c.Delete(uri, c.addHttpHeaders)

	return newResult(resp, err)
}

// issue a Get http request to the resource URI
func (c *Client) GetItem(item *Item) (*Item, error) {
	uri, err := item.uri(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}

	// make an http Get request to the service
	result, err := c.Get(uri, c.addHttpHeaders)

	if err != nil {
		return nil, err
	}

	i, err := item.decode(result)

	defer func() {
		if ferr := result.Body.Close(); ferr != nil {
			err = ferr
		}
	}()

	return i, err
}

// Get a list of items which are linked to the specified item
func (c *Client) GetItemChildren(item *Item) (*ItemList, error) {
	uri, err := item.uriItemChildren(c.conf.BaseURI)
	if err != nil {
		return nil, err
	}

	// make an http Get request to the service
	result, err := c.Get(uri, c.addHttpHeaders)

	if err != nil {
		return nil, err
	}

	list, err := item.decodeList(result)

	defer func() {
		if ferr := result.Body.Close(); ferr != nil {
			err = ferr
		}
	}()

	return list, err
}
