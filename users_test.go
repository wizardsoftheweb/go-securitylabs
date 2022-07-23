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
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
