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

// GET "/api/engagement/time?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d"
// responses pulled from
// https://apidocs.hunter2.com/#get-engagement-time
// I have no idea if these are actually what the API returns
func handlerGetEngagementTime(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetEngagementTime: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"duration":1800000}`))
}
