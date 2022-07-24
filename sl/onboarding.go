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
	GetOnboardingPath = "/onboarding"
)

// GetOnboardingMilestones
// This is the format of the milestones object on the /api/onboarding endpoint
// https://apidocs.hunter2.com/#get-onboarding
type GetOnboardingMilestones struct {
	Total       int `json:"total"`
	Finished    int `json:"finished"`
	NotStarted  int `json:"notStarted"`
	NotSignedUp int `json:"notSignedUp"`
	Started     int `json:"started"`
}

// GetOnboardingUsers
// This is the format of the users object on the /api/onboarding endpoint
// https://apidocs.hunter2.com/#get-onboarding
type GetOnboardingUsers struct {
	Id         string           `json:"id"`
	LastActive *int64           `json:"lastActive"`
	Milestone  string           `json:"milestone"`
	Name       string           `json:"name"`
	Roles      []RolesWithNames `json:"roles"`
}

// GetOnboardingResponse
// This is the full body of the /api/onboarding endpoint
// https://apidocs.hunter2.com/#get-onboarding
type GetOnboardingResponse struct {
	Milestones GetOnboardingMilestones `json:"milestones"`
	Pages      Pages                   `json:"pages"`
	Users      []GetOnboardingUsers    `json:"users"`
}
