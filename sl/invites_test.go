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

// POST /api/invites
// responses pulled from
// https://apidocs.hunter2.com/#post-invites
// I have no idea if these are actually what the API returns
func handlerPostInvites(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerPostInvites: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "email": "test@hunter2.com",
  "senderId": "cb4f412c59ef801d0e6de1c6",
  "roleIds": ["4acc432b734d1d55c318ef58"],
  "admin": false,
}`))
}
