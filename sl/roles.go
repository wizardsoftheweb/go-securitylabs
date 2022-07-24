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

const (
	GetRolesPath         = "/roles"
	GetRolesProgressPath = "/roles/%s/progress"
)

// GetRolesRole
// This is the format of the role object /api/roles endpoint
// https://apidocs.hunter2.com/#get-roles
type GetRolesRole struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Default      bool     `json:"default"`
	Public       bool     `json:"public"`
	Users        []UserId `json:"users"`
	InvitedUsers []UserId `json:"invitedUsers"`
}

// GetRolesResponse
// This is the full body of the /api/roles endpoint
// https://apidocs.hunter2.com/#get-roles
type GetRolesResponse []GetRolesRole
