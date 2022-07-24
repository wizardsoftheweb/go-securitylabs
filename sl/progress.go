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
	GetProgressPath = "/progress"
)

// GetProgressUsersLesson
// This is the format of the lessons object on the users object on the /api/progress endpoint
// https://apidocs.hunter2.com/#get-progress
type GetProgressUsersLesson struct {
	LessonId   string  `json:"lessonId"`
	LessonName string  `json:"lessonName"`
	Module     string  `json:"module"`
	StartTime  int64   `json:"startTime"`
	EndTime    *string `json:"endTime"`
}

// GetProgressUsers
// This is the format of the users object on the /api/progress endpoint
// https://apidocs.hunter2.com/#get-progress
type GetProgressUsers struct {
	Id                        string                   `json:"id"`
	Email                     string                   `json:"email"`
	Name                      string                   `json:"name"`
	LabsCompleted             int                      `json:"labsCompleted"`
	LastActive                int64                    `json:"lastActive"`
	TotalPoints               int                      `json:"totalPoints"`
	LabsStarted               int                      `json:"labsStarted"`
	AccountDisabled           bool                     `json:"accountDisabled"`
	RequiredCompletionPercent int                      `json:"requiredCompletionPercent"`
	PointsRequired            int                      `json:"pointsRequired"`
	PointsPossible            int                      `json:"pointsPossible"`
	Lessons                   []GetProgressUsersLesson `json:"lessons"`
	Roles                     []RoleName               `json:"roles"`
}

// GetProgressResponse
// This is the full body of the /api/progress endpoint
// https://apidocs.hunter2.com/#get-progress
type GetProgressResponse struct {
	NextPage *string            `json:"nextPage"`
	Users    []GetProgressUsers `json:"users"`
}
