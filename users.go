// Copyright 2022 CJ Harries
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vsl

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Roles []string

type Users struct {
	Id         string `json:"id"`
	IsAdmin    bool   `json:"isAdmin"`
	IsDisabled bool   `json:"isDisabled"`
	Email      string `json:"email"`
	Joined     bool   `json:"joined"`
	LastActive int64  `json:"lastActive"`
	Roles      Roles  `json:"roles"`
}

type UsersResponse struct {
	NextPage string  `json:"nextPage"`
	Users    []Users `json:"users"`
}

func (c *Client) GetUsers() ([]Users, error) {
	relativeUrl := &url.URL{Path: "/users"}
	requestUrl := c.Config.BaseUrl.ResolveReference(relativeUrl)
	request, requestGenerationErr := http.NewRequest("GET", requestUrl.String(), nil)
	if nil != requestGenerationErr {
		return nil, requestGenerationErr
	}
	request.Header.Set("Accept", "application/json")
	response, responseErr := c.httpClient.Do(request)
	if nil != responseErr {
		return nil, responseErr
	}
	defer (func() { _ = response.Body.Close() })()
	var responseBody UsersResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseBody)
	if nil != decodeErr {
		return nil, decodeErr
	}
	return responseBody.Users, nil
}
