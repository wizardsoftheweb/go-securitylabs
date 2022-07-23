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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hetiansu5/urlquery"
	"github.com/stretchr/testify/assert"
)

const (
	testMaxPage = 10
)

// /api/users?page=1
// responses pulled from
// https://apidocs.hunter2.com/#get-users
// I have no idea if these are actually what the API returns
func usersTestHandler(w http.ResponseWriter, r *http.Request) {
	var params UsersOptions
	_ = urlquery.Unmarshal([]byte(r.URL.RawQuery), &params)
	var page string
	if params.Page > testMaxPage {
		page = "null"
	} else {
		page = fmt.Sprintf("\"/api/users?page=%d\"", params.Page+1)
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

//// "/api/users/details?campaignIds=5f5f18ff9dad493352660d2a,5f5f19099dad493352660d2b&roleIds=5f5f190f9dad493352660d2c,5f5f19159dad493352660d2d&startTime=1600067874107&endTime=1600067881636&sort=name&sortType=ASC&phrase=Chris&limit=10&page=0"
//// responses pulled from
//// https://apidocs.hunter2.com/#get-users-details
//// I have no idea if these are actually what the API returns
//func usersDetailsTestHandler(w http.ResponseWriter, r *http.Request) {
//	convertedPage, pageConversionErr := strconv.Atoi(r.URL.Query().Get("page"))
//	if pageConversionErr != nil || 0 > convertedPage {
//		convertedPage = 0
//	}
//	//var nextPage, previousPage string
//	//if convertedPage > testMaxPage {
//	//	nextPage = "null"
//	//} else {
//	//
//	//}
//}

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
