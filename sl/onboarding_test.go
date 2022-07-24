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

// GET "/api/onboarding?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&sortType=ASC&phrase=Chris&limit=10&page=0"
// responses pulled from
// https://apidocs.hunter2.com/#get-onboarding
// I have no idea if these are actually what the API returns
func handlerGetOnboarding(w http.ResponseWriter, r *http.Request) {
	// TODO: handlerGetOnboarding: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "milestones": {
    "total": 1,
    "finished": 1,
    "notStarted": 0,
    "notSignedUp": 0,
    "started": 0
  },
  "pages": {
    "current": 1,
    "previous": 0,
    "next": 2,
    "limit": 10,
    "total": 25,
    "currentUrl": "/api/onboarding?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&phrase=Chris&limit=10&page=0",
    "nextUrl": "/api/onboarding?campaignIds=6f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&phrase=Chris&limit=10&page=1",
    "previousUrl": "/api/onboarding?campaignIds=6f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&phrase=Chris&limit=10&page=0"
  },
  "users": [{
    "id": "5f5f18439dad493352660d28",
    "lastActive": 1600067617569,
    "name": "Chris Traeger",
    "milestone": "finished",
    "roles": [{
      "id": "5f5f184d9dad493352660d29",
      "name": "Security"
    }]
  }]
}`))
}
