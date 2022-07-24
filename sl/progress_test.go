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

// GET /api/progress?page=0
// responses pulled from
// https://apidocs.hunter2.com/#get-progress
// I have no idea if these are actually what the API returns
func handlerGetProgress(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetProgress: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "nextPage": "/api/progress?page=1",
  "users": [{
    "id": "3bd68695e165af6ced227afc",
    "email": "test@hunter2.com",
    "name": "Test User",
    "labsCompleted": 1,
    "lastActive": "3/6/2019",
    "totalPoints": 10,
    "labsStarted": 2,
    "disabled": false,
    "requiredCompletionPercent": 14,
    "pointsRequired" : 10,
    "pointsPossible" : 30,
    "lessons": [{
      "lessonId": "5a5999d4ca50092ec5345ec4",
      "lessonName": "Own the database",
      "module": "OWASP #1: Injection",
      "startTime": 1551673334046,
      "endTime": 1551673334047,
      "startRating": 1,
      "endRating": 2,
      "lastStepReached": 1,
      "status": "Finished"
    }]
    "roles": [
      "Developers"
    ]
  }]
}`))
}
