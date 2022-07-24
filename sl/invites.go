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
	PostInvitePath = "/invite"
)

// PostInvites
// This is the format of the request body on the /api/invites endpoint
// https://apidocs.hunter2.com/#post-invites
// TODO: Can name be sent as well?
type PostInvites struct {
	Email     string   `json:"email"`
	SenderId  UserId   `json:"senderId"`
	RoleIds   []RoleId `json:"roleIds"`
	Admin     bool     `json:"admin"`
	SendEmail bool     `json:"sendEmail"`
}

// PostInvitesResponse
// This is the format of the response body on the /api/invites endpoint
// https://apidocs.hunter2.com/#post-invites
type PostInvitesResponse struct {
	Email    string   `json:"email"`
	Admin    bool     `json:"admin"`
	RoleIds  []RoleId `json:"roleIds"`
	SenderId UserId   `json:"senderId"`
}
