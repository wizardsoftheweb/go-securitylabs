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
	GetLessonsPath = "/api/lessons"
)

// GetLessonsLesson
// This is the format of the lesson object on the /api/lessons endpoint
// https://apidocs.hunter2.com/#get-lessons
type GetLessonsLesson struct {
	Id     string     `json:"id"`
	Title  string     `json:"title"`
	Module string     `json:"module"`
	Points int        `json:"points"`
	Roles  []RoleName `json:"roles"`
}

// GetLessonsResponse
// This is the full body of the /api/lessons endpoint
// https://apidocs.hunter2.com/#get-lessons
type GetLessonsResponse struct {
	NextPage *string            `json:"nextPage"`
	Lessons  []GetLessonsLesson `json:"lessons"`
}
