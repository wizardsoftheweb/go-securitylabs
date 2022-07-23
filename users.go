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

// UsersUsersRoles
// This is a list of roles attached to a user
// https://apidocs.hunter2.com/#get-users
type UsersUsersRoles []string

// UsersUsers
// These are all the properties of a user from the /api/user endpoint
// https://apidocs.hunter2.com/#get-users
type UsersUsers struct {
	Id         string          `json:"id"`
	IsAdmin    bool            `json:"isAdmin"`
	IsDisabled bool            `json:"isDisabled"`
	Email      string          `json:"email"`
	Joined     bool            `json:"joined"`
	LastActive int64           `json:"lastActive"`
	Roles      UsersUsersRoles `json:"roles"`
}

// UsersResponse
// This is the full body of the /api/user endpoint
// https://apidocs.hunter2.com/#get-users
type UsersResponse struct {
	NextPage string       `json:"nextPage"`
	Users    []UsersUsers `json:"users"`
}

// UsersOptions
// These are the query params for the /api/users endpoint
// https://apidocs.hunter2.com/#get-users
type UsersOptions struct {
	Page *int `query:"page"`
}

func (c *Client) GetUsers() ([]UsersUsers, error) {
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

// UsersDetailsPages
// This is the format of the pages object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type UsersDetailsPages struct {
	Current     int    `json:"current"`
	Previous    int    `json:"previous"`
	Next        int    `json:"next"`
	Limit       int    `json:"limit"`
	Total       int    `json:"total"`
	CurrentUrl  string `json:"currentUrl"`
	NextUrl     string `json:"nextUrl"`
	PreviousUrl string `json:"previousUrl"`
}

// UsersDetailsUsersRoles
// This is the format of the roles object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type UsersDetailsUsersRoles struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// UsersDetailsUsers
// This is the format of the users object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type UsersDetailsUsers struct {
	Id                       string                   `json:"id"`
	LastActive               int64                    `json:"lastActive"`
	LabsCompleted            int                      `json:"labsCompleted"`
	PercentRequiredCompleted float64                  `json:"percentRequiredCompleted"`
	Points                   int                      `json:"points"`
	Name                     string                   `json:"name"`
	Roles                    []UsersDetailsUsersRoles `json:"roles"`
}

// UsersDetailsResponse
// This is the full body of the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type UsersDetailsResponse struct {
	Pages UsersDetailsPages   `json:"pages"`
	Users []UsersDetailsUsers `json:"users"`
}

// UsersDetailsOptions
// These are the query params for the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type UsersDetailsOptions struct {
	CampaignIds []string `query:"campaignIds"`
	EndTime     *int64   `query:"endTime"`
	Limit       *int     `query:"limit"`
	Page        *int     `query:"page"`
	Phrase      *string  `query:"phrase"`
	StartTime   *int64   `query:"startTime"`
	RoleIds     []string `query:"roleIds"`
	Sort        *string  `query:"sort"`
	SortType    *string  `query:"sortType"`
}
