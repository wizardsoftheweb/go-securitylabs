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
	GetLessonsPath         = "/api/lessons"
	GetLessonsProgressPath = "/api/lessons/%s/progress"
	GetLessonsSearchPath   = "/api/lessons/search"
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

// GetLessonsProgressUser
// This is the format of the user object on the /api/lessons/{lessonId}/progress endpoint
// https://apidocs.hunter2.com/#get-lessons-progress
type GetLessonsProgressUser struct {
	Id              string `json:"id"`
	StartTime       *int64 `json:"startTime"`
	EndTime         *int64 `json:"endTime"`
	StartRating     *int   `json:"startRating"`
	EndRating       *int   `json:"endRating"`
	LastStepReached *int   `json:"lastStepReached"`
	Status          string `json:"status"`
}

// GetLessonsProgressResponse
// This is the full body of the /api/lessons/{lessonId}/progress endpoint
// https://apidocs.hunter2.com/#get-lessons-progress
type GetLessonsProgressResponse struct {
	NextPage *string                  `json:"nextPage"`
	Id       string                   `json:"id"`
	Title    string                   `json:"title"`
	Module   string                   `json:"module"`
	Points   int                      `json:"points"`
	Roles    []RoleName               `json:"roles"`
	Users    []GetLessonsProgressUser `json:"users"`
}

// GetLessonsSearchLesson
// This is the format of the lesson object on the /api/lessons/search endpoint
// https://apidocs.hunter2.com/#get-lesson-by-topic
type GetLessonsSearchLesson struct {
	Challenge   bool     `json:"challenge"`
	Description string   `json:"description"`
	Stack       string   `json:"stack"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Topic       string   `json:"topic"`
	Url         string   `json:"url"`
}

// GetLessonsSearchResponse
// This is the full body of the /api/lessons/search endpoint
// https://apidocs.hunter2.com/#get-lesson-by-topic
type GetLessonsSearchResponse struct {
	Pages   Pages                    `json:"pages"`
	Lessons []GetLessonsSearchLesson `json:"lessons"`
}

// GetLessonsSearchOptions
// These are the query params on the /api/lessons/search endpoint
// https://apidocs.hunter2.com/#get-lesson-by-topic
type GetLessonsSearchOptions struct {
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
	Phrase string `query:"phrase"`
}
