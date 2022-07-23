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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestNewClient(t *testing.T) {
	client := NewClient(nil, nil)
	assert.NotNilf(t, client, "Client should not be nil")
	assert.Equalf(t, productionUrl, client.BaseUrl, "BaseUrl should be set to productionUrl")
	clientWithUrl := NewClient(&url.URL{}, nil)
	assert.NotNilf(t, clientWithUrl, "Client should not be nil")
	assert.NotEqualf(t, productionUrl, clientWithUrl.BaseUrl, "BaseUrl should not be set to productionUrl")
	clientWithHttpClient := NewClient(nil, &http.Client{Timeout: time.Duration(1) * time.Second})
	assert.NotNilf(t, clientWithHttpClient, "Client should not be nil")
	assert.NotEqualf(t, clientWithHttpClient.httpClient, http.DefaultClient, "httpClient should not be set to http.DefaultClient")
}

type ClientNewRequestSuite struct {
	suite.Suite
	server    *httptest.Server
	serverUrl *url.URL
	client    *Client
}

func TestClientNewRequestSuite(t *testing.T) {
	suite.Run(t, new(ClientNewRequestSuite))
}

func (suite *ClientNewRequestSuite) SetupTest() {
	mux := http.NewServeMux()
	suite.server = httptest.NewServer(mux)
	defer suite.server.Close()
	suite.serverUrl, _ = url.Parse(suite.server.URL)
	suite.client = NewClient(suite.serverUrl, nil)
}

func (suite *ClientNewRequestSuite) TestClient_newRequest_Empty() {
	request, err := suite.client.newRequest(http.MethodGet, "/", nil)
	suite.Nilf(err, "Error should be nil")
	suite.NotNilf(request, "Request should not be nil")
	suite.Equalf(http.MethodGet, request.Method, "Method should be set to http.MethodGet")
	suite.Equalf(suite.server.URL+"/", request.URL.String(), "URL should be set to testServer.URL")
	suite.Nilf(request.Body, "Body should be nil")
}

func (suite *ClientNewRequestSuite) TestClient_newRequest_WithBody() {
	body := struct {
		Name string `json:"name"`
	}{Name: "test"}
	request, err := suite.client.newRequest(http.MethodPost, "/", body)
	suite.Nilf(err, "Error should be nil")
	suite.NotNilf(request, "Request should not be nil")
	suite.Equalf(http.MethodPost, request.Method, "Method should be set to http.MethodPost")
	suite.Equalf(suite.server.URL+"/", request.URL.String(), "URL should be set to testServer.URL")
	suite.NotNilf(request.Body, "Body should not be nil")
	suite.Equalf("application/json", request.Header.Get("Content-Type"), "Content-Type should be set to application/json")
	suite.Equalf("application/json", request.Header.Get("Accept"), "Accept should be set to application/json")
	parsedBody, _ := ioutil.ReadAll(request.Body)
	suite.Equalf("{\"name\":\"test\"}\n", string(parsedBody), "Body should be set to `{\"name\":\"test\"}\n`")
}

func (suite *ClientNewRequestSuite) TestClient_newRequest_WithBodyBuildError() {
	body := make(chan int)
	request, err := suite.client.newRequest(http.MethodPost, "/", body)
	suite.Nilf(request, "Request should be nil")
	suite.NotNilf(err, "Error should not be nil")
}

//type ClientDoResponse struct {
//	Message string `json:"message"`
//}
//
//func TestClient_Do(t *testing.T) {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
//	})
//	server := httptest.NewServer(mux)
//	defer server.Close()
//	client := NewClient(&ClientConfig{BaseUrl: &url.URL{Host: server.URL}}, nil)
//	assert.NotNilf(t, client, "Client should not be nil")
//	var doResponse *ClientDoResponse
//	response, err := client.do(context.Background(), &http.Request{}, doResponse)
//	fmt.Printf("%+v\n", err)
//	fmt.Printf("%+v\n", response)
//	assert.Nilf(t, err, "Error should be nil")
//	assert.Equalf(t, http.StatusOK, response.StatusCode, "StatusCode should be 200")
//	assert.NotNilf(t, doResponse, "doResponse should not be nil")
//	assert.Equalf(t, "ok", doResponse.Message, "Message should be ok")
//}
