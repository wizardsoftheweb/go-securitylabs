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

package sl

// Pages
// This is the format of the pages object several endpoints
type Pages struct {
	Current     int     `json:"current"`
	Previous    *int    `json:"previous"`
	Next        *int    `json:"next"`
	Limit       int     `json:"limit"`
	Total       int     `json:"total"`
	CurrentUrl  string  `json:"currentUrl"`
	NextUrl     *string `json:"nextUrl"`
	PreviousUrl *string `json:"previousUrl"`
}

// UsersDetailsOptions
// These are the query params to limit users on several endpoints
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

// RolesWithNames
// This is a list of roles attached to a user on several endpoints
type RolesWithNames struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// UsersWithActivity
// This is the format of the users object on several endpoints
type UsersWithActivity struct {
	Id         string           `json:"id"`
	LastActive *int64           `json:"lastActive"`
	Milestone  string           `json:"milestone"`
	Name       string           `json:"name"`
	Roles      []RolesWithNames `json:"roles"`
}

// PageOptions
// This handles options for endpoints that only allow page as a param
type PageOptions struct {
	Page *int `query:"page"`
}

// RoleName
// It's intended to be used when roles is the names, not the IDs
type RoleName string

// RoleId
// It's intended to be used when roles is the IDs, not the names
type RoleId string
