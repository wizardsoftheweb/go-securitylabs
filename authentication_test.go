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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testKey            = "proper"
	testSecret         = "secret"
	testRotatedSecret  = "rotated"
	testDisabledSecret = "disabled"
)

// Authorization responses pulled from
// https://apidocs.hunter2.com/#authentication
// I have no idea if these are actually what the API returns
func authenticationTestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authKey, secret string
		for key, value := range r.Header {
			if strings.ToLower(key) == "auth" {
				exploded := strings.Split(value[0], ":")
				authKey, secret = exploded[0], exploded[1]
				break
			}
		}
		fmt.Println("authKey:", authKey, "secret:", secret)
		if "" == authKey || "" == secret {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte("{\"message\":\"ApiCredential invalid\"}"))
			return
		}
		if testKey != authKey {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("{\"message\":\"ApiCredential missing\"}"))
			return
		} else {
			if testSecret == secret {
				next.ServeHTTP(w, r)
			}
			if testDisabledSecret == secret {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("{\"message\":\"disabled\"}"))
				return
			}
			if testRotatedSecret == secret {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("{\"message\":\"rotated\"}"))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("{\"message\":\"ApiCredential missing\"}"))
	})
}

type AuthenticationTestSuite struct {
	suite.Suite
	server             *httptest.Server
	serverUrl          *url.URL
	client             *Client
	existingAuthKey    string
	existingAuthSecret string
}

func TestAuthenticationTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}

func (suite *AuthenticationTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.Handle("/ok", authenticationTestMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
	})))
	suite.server = httptest.NewServer(mux)
	suite.serverUrl, _ = url.Parse(suite.server.URL)
	suite.client = NewClient(suite.serverUrl, nil)
	suite.existingAuthKey = os.Getenv(EnvVslAuthKey)
	suite.existingAuthSecret = os.Getenv(EnvVslAuthSecret)
	suite.client.AuthKey = testKey
	suite.client.AuthSecret = testSecret
	_ = os.Unsetenv(EnvVslAuthKey)
	_ = os.Unsetenv(EnvVslAuthSecret)
}

func (suite *AuthenticationTestSuite) TearDownTest() {
	suite.server.Close()
	if "" != suite.existingAuthKey {
		_ = os.Setenv(EnvVslAuthKey, suite.existingAuthKey)
	} else {
		_ = os.Unsetenv(EnvVslAuthKey)
	}
	if "" != suite.existingAuthSecret {
		_ = os.Setenv(EnvVslAuthSecret, suite.existingAuthSecret)
	} else {
		_ = os.Unsetenv(EnvVslAuthSecret)
	}
}

func (suite *AuthenticationTestSuite) TestClient_SetAuthFromEnvironment() {
	client := NewClient(suite.serverUrl, nil)
	suite.NotEqualf(testKey, client.AuthKey, "AuthKey should be set from environment")
	suite.NotEqualf(testSecret, client.AuthSecret, "AuthSecret should be set from environment")
	noAuthKeyError := client.SetAuthFromEnvironment()
	suite.NotNilf(noAuthKeyError, "Should return error if no AuthKey is set")
	_ = os.Setenv(EnvVslAuthKey, testKey)
	noAuthSecretError := client.SetAuthFromEnvironment()
	suite.NotNilf(noAuthSecretError, "Should return error if no AuthSecret is set")
	_ = os.Setenv(EnvVslAuthSecret, testSecret)
	noError := client.SetAuthFromEnvironment()
	suite.Nilf(noError, "SetAuthFromEnvironment should not return an error")
	suite.Equalf(testKey, client.AuthKey, "AuthKey should be set from environment")
	suite.Equalf(testSecret, client.AuthSecret, "AuthSecret should be set from environment")
}

func (suite *AuthenticationTestSuite) TestClient_SetAuth() {
	client := NewClient(suite.serverUrl, nil)
	suite.NotEqualf(testKey, client.AuthKey, "AuthKey should be set from fnc")
	suite.NotEqualf(testSecret, client.AuthSecret, "AuthSecret should be set from fnc")
	client.SetAuth(testKey, testSecret)
	suite.Equalf(testKey, client.AuthKey, "AuthKey should be set from fnc")
	suite.Equalf(testSecret, client.AuthSecret, "AuthSecret should be set from fnc")
}

type ClientAuthResponse struct {
	Message string `json:"message"`
}

func (suite *AuthenticationTestSuite) TestClient_Req_Success() {
	request, requestGenerationError := suite.client.newRequest("GET", "/ok", nil)
	suite.Nilf(requestGenerationError, "Should not return error when generating request")
	var authResponse *ClientAuthResponse
	response, responseError := suite.client.do(context.Background(), request, &authResponse)
	suite.Nilf(responseError, "Should not return error when making request")
	suite.Equalf(http.StatusOK, response.StatusCode, "Should return status code 200")
}

func (suite *AuthenticationTestSuite) TestClient_Req_NoAuth() {
	suite.client.AuthKey = ""
	suite.client.AuthSecret = ""
	request, requestGenerationErr := suite.client.newRequest("GET", "/ok", nil)
	suite.Nilf(requestGenerationErr, "Should not return error when creating request")
	var authResponse *ClientAuthResponse
	response, responseError := suite.client.do(context.Background(), request, &authResponse)
	suite.Nilf(responseError, "Should not return error when making request")
	suite.Equalf(http.StatusForbidden, response.StatusCode, "Should return 403 status code")
}

func (suite *AuthenticationTestSuite) TestClient_Req_AuthMissing() {
	suite.client.AuthKey = "qqq"
	request, requestGenerationErr := suite.client.newRequest("GET", "/ok", nil)
	suite.Nilf(requestGenerationErr, "Should not return error when creating request")
	var authResponse *ClientAuthResponse
	response, responseError := suite.client.do(context.Background(), request, &authResponse)
	suite.Nilf(responseError, "Should not return error when making request")
	suite.Equalf(http.StatusNotFound, response.StatusCode, "Should return 404 status code")
}

func (suite *AuthenticationTestSuite) TestClient_Req_AuthRotated() {
	suite.client.AuthSecret = testRotatedSecret
	request, requestGenerationErr := suite.client.newRequest("GET", "/ok", nil)
	suite.Nilf(requestGenerationErr, "Should not return error when creating request")
	var authResponse *ClientAuthResponse
	response, responseError := suite.client.do(context.Background(), request, &authResponse)
	suite.Nilf(responseError, "Should not return error when making request")
	suite.Equalf(http.StatusUnauthorized, response.StatusCode, "Should return 401 status code")
}

func (suite *AuthenticationTestSuite) TestClient_Req_AuthDisabled() {
	suite.client.AuthSecret = testDisabledSecret
	request, requestGenerationErr := suite.client.newRequest("GET", "/ok", nil)
	suite.Nilf(requestGenerationErr, "Should not return error when creating request")
	var authResponse *ClientAuthResponse
	response, responseError := suite.client.do(context.Background(), request, &authResponse)
	suite.Nilf(responseError, "Should not return error when making request")
	suite.Equalf(http.StatusUnauthorized, response.StatusCode, "Should return 401 status code")
}
