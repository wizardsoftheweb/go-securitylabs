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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hetiansu5/urlquery"
)

const (
	testMaxPage           = 10
	testMissingCampaignId = "qqq"
	testMissingRoleId     = "qqq"
	testLimitMin          = 0
	testLimitMax          = 50
	testPhraseMax         = 50
	testTotal             = 100
	testNonexistentUserId = "5f5f18439qqq493352660d28"
	testNameMax           = 25
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
	testExistingUserIds = []string{
		"3bd68695e165af6ced227afc",
		"5f5f18439dad493352660d28",
		testNonexistentUserId,
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

// GET /api/users?page=1
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

// GET /api/users/details?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&sortType=ASC&phrase=Chris&limit=10&page=0
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

// GET /api/users/:id/progress
// responses pulled from
// https://apidocs.hunter2.com/#get-user-progress
// I have no idea if these are actually what the API returns
func handlerGetUserProgress(w http.ResponseWriter, r *http.Request) {
	userId := strings.Replace(r.URL.RequestURI(), "/api/user/", "", 1)
	// TODO: How to test ID validity?
	if 24 != len(userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User id is invalid\"}"))
		return
	}
	if !listContains(testExistingUserIds, userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User not found\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "pointsRequired": 10,
  "pointsPossible": 30,
  "lessons": [{
    "module": "OWASP #1: Injection",
    "lessonId": "5a5999d4ca50092ec5345ec4",
    "lessonName": "Own the database",
    "lastVisited": "3/4/2019",
    "status": "Started",
    "minutes": 1216.9,
    "points": 0,
    "startRating": 1,
    "endRating": 2
  }]
}`))
}

// PUT /api/users/:id
// responses pulled from
// https://apidocs.hunter2.com/#put-user
// I have no idea if these are actually what the API returns
func handlerPutUser(w http.ResponseWriter, r *http.Request) {
	var user PutUser
	parseError := json.NewDecoder(r.Body).Decode(&user)
	// TODO: is this the correct error?
	if nil != parseError {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"Invalid request body\"}"))
		return
	}
	userId := strings.Replace(r.URL.RequestURI(), "/api/user/", "", 1)
	// TODO: How to test ID validity?
	if 24 != len(userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User id is invalid\"}"))
		return
	}
	if !listContains(testExistingUserIds, userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User not found\"}"))
		return
	}
	// TODO: How to check email validity?
	if "" == user.Email {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"Email is invalid\"}"))
		return
	}
	if testNameMax < len(user.Name) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"Name must be less than 25 characters\"}"))
		return
	}
	// TODO: How to check role ids not being an array?
	// TODO: Does the endpoint actually check valid role ids?
	for _, role := range user.RoleIds {
		if !listContains(testRoleIds, role) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"message\":\"Invalid roleIds\"}"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(user)
	_, _ = w.Write(responseBody)
}

// DELETE /api/users/:id
// responses pulled from
// https://apidocs.hunter2.com/#delete-user
// I have no idea if these are actually what the API returns
func handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := strings.Replace(r.URL.RequestURI(), "/api/user/", "", 1)
	// TODO: How to test ID validity?
	if 24 != len(userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User id is invalid\"}"))
		return
	}
	if !listContains(testExistingUserIds, userId) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("{\"message\":\"User not found\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{}"))
}

type UsersTestSuite struct {
	suite.Suite
	server    *httptest.Server
	serverUrl *url.URL
	client    *Client
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (suite *UsersTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.Handle("/users", http.HandlerFunc(handlerGetUsers))
	suite.server = httptest.NewServer(mux)
	suite.serverUrl, _ = url.Parse(suite.server.URL)
	suite.client = NewClient(suite.serverUrl, nil)
}

func (suite *UsersTestSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *UsersTestSuite) TestClient_GetUsers() {
	users, usersErr := suite.client.GetUsers(context.Background(), nil)
	suite.Nilf(usersErr, "GetUsers() should not return an error")
	suite.Truef(len(users) > 0, "GetUsers() should return at least one user")
}
