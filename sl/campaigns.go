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
	GetCampaignsProgressPath = "/campaigns/progress"
)

// GetCampaignsProgressMilestones
// This is the format of the milestones object on several components of the /api/campaigns/progress endpoint
// https://apidocs.hunter2.com/#get-campaigns-progress
type GetCampaignsProgressMilestones struct {
	Total      int `json:"total"`
	Finished   int `json:"finished"`
	NotStarted int `json:"notStarted"`
	Started    int `json:"started"`
}

// GetCampaignsProgressAssignmentsLab
// This is the format of the assignmentsLabs object on the assignments object on /api/campaigns/progress endpoint
// https://apidocs.hunter2.com/#get-campaigns-progress
type GetCampaignsProgressAssignmentsLab struct {
	Title      string                         `json:"title"`
	Milestones GetCampaignsProgressMilestones `json:"milestones"`
}

// GetCampaignsProgressAssignment
// This is the format of the assignments object on the /api/campaigns/progress endpoint
// https://apidocs.hunter2.com/#get-campaigns-progress
type GetCampaignsProgressAssignment struct {
	Title string                               `json:"title"`
	Labs  []GetCampaignsProgressAssignmentsLab `json:"labs"`
}

// GetCampaignsProgressResponse
// This is the full body of the /api/campaigns/progress endpoint
// https://apidocs.hunter2.com/#get-campaigns-progress
type GetCampaignsProgressResponse struct {
	Milestones  GetCampaignsProgressMilestones   `json:"milestones"`
	Assignments []GetCampaignsProgressAssignment `json:"assignments"`
	Pages       Pages                            `json:"pages"`
	Users       []UserWithActivity               `json:"users"`
}

// GetCampaignsProgessOptions
// This is the format of the options object on the /api/campaigns/progress endpoint
// https://apidocs.hunter2.com/#get-campaigns-progress
type GetCampaignsProgessOptions struct {
	AssignmentIds       []string `query:"assignmentIds"`
	CampaignIds         []string `query:"campaignIds"`
	IncludeOptionalLabs bool     `query:"includeOptionalLabs"`
	Limit               int      `query:"limit"`
	Page                int      `query:"page"`
	Phrase              string   `query:"phrase"`
	RoleIds             []string `query:"roleIds"`
	Sort                string   `query:"sort"`
	SortType            string   `query:"sortType"`
}
