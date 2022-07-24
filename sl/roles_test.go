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

// GET /api/roles
// responses pulled from
// https://apidocs.hunter2.com/#get-roles
// I have no idea if these are actually what the API returns
func handlerGetRoles(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetRoles: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`[{
  "id": "4bd68695e165af6ced227afz",
  "name": "Developers",
  "default": true,
  "public": false,
  "users": [
    "3bd68695e165af6ced227afc"
  ],
  "invitedUsers": [
    "5ce34b1b0e5f930030014122"
  ]
}]`))
}

// GET /api/roles/4bd68695e165af6ced227afz/progress?page=0
// responses pulled from
// https://apidocs.hunter2.com/#get-role-progress
// I have no idea if these are actually what the API returns
func handlerGetRoleProgress(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetRoleProgress: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// TODO: Verify this isn't actually an array of these objects like the docs suggest
	_, _ = w.Write([]byte(`{
  "nextPage": "/api/roles/4bd68695e165af6ced227afz/progress?page=1",
  "id": "4bd68695e165af6ced227afz",
  "name": "Developers",
  "users": [{
    "id": "3bd68695e165af6ced227afc",
    "email": "test@hunter2.com",
    "name": "Test User",
    "percentComplete": 50,
    "percentRequiredComplete": 100
  }]
}`))
}
