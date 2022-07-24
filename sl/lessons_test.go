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

import "net/http"

// GET /api/lessons?page=0
// responses pulled from
// https://apidocs.hunter2.com/#get-lessons
// I have no idea if these are actually what the API returns
func handlerGetLessons(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetLessons: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "nextPage": "/api/lessons?page=1",
  "lessons": [{
    "id": "5a5999d4ca50092ec5345ec4",
    "title": "Own the database",
    "module": "OWASP #1: Injection",
    "points": 20,
    "roles": [
      "Developers"
    ]
  }]
}`))
}

// GET /api/lessons/:id/progress?page=0
// responses pulled from
// https://apidocs.hunter2.com/#get-lesson-progress
// I have no idea if these are actually what the API returns
func handlerGetLessonProgress(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetLessonProgress: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "nextPage": "/api/lessons/5a5999d4ca50092ec5345ec4/progress?page=1",
  "id": "5a5999d4ca50092ec5345ec4",
  "title": "Own the database",
  "module": "OWASP #1: Injection",
  "points": 30,
  "roles": [
    "Developers"
  ],
  "users": [{
    "id": "3bd68695e165af6ced227afc",
    "startTime": 1551673334046,
    "endTime": 1551673334047,
    "startRating": 1,
    "endRating": 4,
    "lastStepReached": 1,
    "status": "Started"
  }]
}`))
}
