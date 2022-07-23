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

package vsl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hetiansu5/urlquery"
	"github.com/stretchr/testify/assert"
)

const (
	testMaxPage           = 10
	testMissingCampaignId = "qqq"
	testMissingRoleId     = "qqq"
	testLimitMin          = 0
	testLimitMax          = 50
	testPhraseMax         = 50
	testTotal             = 100
)

var (
	testCampaignIds = []string{
		"5f5f18ff9dad493352660d2a",
		"5f5f19099dad493352660d2b",
		testMissingCampaignId,
	}
	testRoleIds = []string{
		"5f5f190f9dad493352660d2c",
		"5f5f19159dad493352660d2d",
		testMissingRoleId,
	}
	testValidSorts = []string{
		"name",
		"role",
		"milestone",
		"lastActive",
	}
	testValidSortTypes = []string{
		"ASC",
		"DESC",
	}
)

// Helper function to validate test input
func listContains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

// /api/users?page=1
// responses pulled from
// https://apidocs.hunter2.com/#get-users
// I have no idea if these are actually what the API returns
func handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	var params GetUsersOptions
	_ = urlquery.Unmarshal([]byte(r.URL.RawQuery), &params)
	var page string
	if nil == params.Page || *params.Page > testMaxPage {
		page = "null"
	} else {
		page = fmt.Sprintf("\"/api/users?page=%d\"", *params.Page+1)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(`{
  "nextPage" : %s,
  "users": [{
    "id": "3bd68695e165af6ced227afc",
    "isAdmin": true,
    "isDisabled": false,
    "email": "developer@hunter2.com",
    "joined": true,
    "lastActive": 1557981546394,
    "roles": [
      "Developers"
    ]
  }]
}`, page)))
}

// "/api/users/details?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&sortType=ASC&phrase=Chris&limit=10&page=0"
// responses pulled from
// https://apidocs.hunter2.com/#get-users-details
// I have no idea if these are actually what the API returns
func usersGetUsersDetailsTestHandler(w http.ResponseWriter, r *http.Request) {
	var params GetUsersDetailsOptions
	_ = urlquery.Unmarshal([]byte(r.URL.RawQuery), &params)
	for _, campaignId := range params.CampaignIds {
		if !listContains(testCampaignIds, campaignId) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"Invalid campaign id: String\"}"))
			return
		}
		if testMissingCampaignId == campaignId {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("{\"message\":\"Campaign(s) not found by id(s): String\"}"))
			return
		}
	}
	for _, roleId := range params.RoleIds {
		if !listContains(testRoleIds, roleId) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"Invalid role id: String\"}"))
			return
		}
		if testMissingRoleId == roleId {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("{\"message\":\"Role(s) not found by id(s): String\"}"))
			return
		}
	}
	if nil != params.StartTime && nil != params.EndTime {
		// TODO: which error is thrown here?
		if *params.StartTime > *params.EndTime {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"startTime is invalid\"}"))
			return
		}
		// TODO: Are there other invalid start times?
		// TODO: Are there other invalid end times?
	} else {
		if nil == params.StartTime && nil != params.EndTime {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"startTime is undefined\"}"))
			return
		}
		if nil != params.StartTime && nil == params.EndTime {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"endTime is undefined\"}"))
			return
		}
	}
	if nil != params.Sort && !listContains(testValidSorts, *params.Sort) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"sort is invalid\"}"))
		return
	}
	if nil != params.SortType && !listContains(testValidSortTypes, *params.SortType) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"sortType is invalid\"}"))
		return
	}
	if nil != params.Limit && (*params.Limit < testLimitMin || *params.Limit > testLimitMax) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"limit is invalid\"}"))
		return
	}
	if nil != params.Phrase && len(*params.Phrase) > testPhraseMax {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"phrase is invalid\"}"))
		return
	}
	var nextPage, currentPage, previousPage *int
	if nil == params.Page || 0 == *params.Page {
		currentPage = new(int)
		*currentPage = 0
		nextPage = new(int)
		*nextPage = *params.Page + 1
	} else if *params.Page > testMaxPage {
		currentPage = new(int)
		*currentPage = *params.Page
		previousPage = new(int)
		*previousPage = *params.Page - 1
	} else {
		currentPage = new(int)
		*currentPage = *params.Page
		nextPage = new(int)
		*nextPage = *params.Page + 1
		previousPage = new(int)
		*previousPage = *params.Page - 1
	}
	w.WriteHeader(http.StatusOK)
	pages := GetUsersDetailsPages{
		Current:  *params.Page,
		Previous: previousPage,
		Next:     nextPage,
		Limit:    *params.Limit,
		Total:    testTotal,
	}
	currentParamsBytes, _ := urlquery.Marshal(params)
	currentParams := string(currentParamsBytes)
	if 0 < len(currentParams) {
		currentUrl := fmt.Sprintf("/api/users/details?%s", currentParams)
		pages.CurrentUrl = currentUrl
	} else {
		pages.CurrentUrl = "/api/users/details"
	}
	params.Page = nextPage
	nextParamsBytes, _ := urlquery.Marshal(params)
	nextParams := string(nextParamsBytes)
	if 0 < len(nextParams) {
		nextUrl := fmt.Sprintf("/api/users/details?%s", nextParams)
		pages.NextUrl = &nextUrl
	} else {
		pages.NextUrl = new(string)
		*pages.NextUrl = "/api/users/details"
	}
	params.Page = previousPage
	previousParamsBytes, _ := urlquery.Marshal(params)
	previousParams := string(previousParamsBytes)
	if 0 < len(previousParams) {
		previousUrl := fmt.Sprintf("/api/users/details?%s", previousParams)
		pages.PreviousUrl = &previousUrl
	} else {
		pages.PreviousUrl = new(string)
		*pages.PreviousUrl = "/api/users/details"
	}
	pagesBytes, _ := json.Marshal(pages)
	_, _ = w.Write([]byte(fmt.Sprintf(`{
  "pages": %s,
  "users": [{
    "id": "5f5f18439dad493352660d28",
    "lastActive": 1600067617569,
    "name": "Chris Traeger",
    "labsCompleted": 1,
    "percentRequiredComplete": 50,
    "points": 10,
    "roles": [{
      "id": "5f5f184d9dad493352660d29",
      "name": "Security"
    }]
  }]
}`, string(pagesBytes))))
}

func TestClient_GetUsers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
  "nextPage" : "/api/users?page=1",
  "users": [{
    "id": "3bd68695e165af6ced227afc",
    "isAdmin": true,
    "isDisabled": false,
    "email": "developer@hunter2.com",
    "joined": true,
    "lastActive": 1557981546394,
    "roles": [
      "Developers"
    ]
  }]
}
`))
	})
	testServer := httptest.NewServer(mux)
	defer testServer.Close()
	testServerUrl, _ := url.Parse(testServer.URL)
	client := NewClient(&ClientConfig{BaseUrl: testServerUrl}, nil)
	assert.Equalf(t, testServerUrl, client.Config.BaseUrl, "BaseUrl should be set")
	users, err := client.GetUsers()
	assert.Nilf(t, err, "Response error should be nil")
	assert.Equal(t, 1, len(users))
}
