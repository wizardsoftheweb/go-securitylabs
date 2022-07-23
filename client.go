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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Get the latest prod URL: https://apidocs.hunter2.com/#production
var productionUrl, _ = url.Parse("https://securitylabs.veracode.com/api/")

type Client struct {
	BaseUrl    *url.URL
	httpClient *http.Client
}

func NewClient(baseUrl *url.URL, httpClient *http.Client) *Client {
	var newHttpClient *http.Client
	var newBaseUrl *url.URL
	if nil == baseUrl {
		newBaseUrl = productionUrl
	} else {
		newBaseUrl = baseUrl
	}
	if nil == httpClient {
		newHttpClient = http.DefaultClient
	} else {
		newHttpClient = httpClient
	}
	return &Client{
		BaseUrl:    newBaseUrl,
		httpClient: newHttpClient,
	}
}

// Create a request to the API with the given method, path, and body
// https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	relativeUrl := &url.URL{Path: path}
	fullUrl := c.BaseUrl.ResolveReference(relativeUrl)
	var buffer io.ReadWriter
	if nil != body {
		buffer = new(bytes.Buffer)
		jsonEncoderError := json.NewEncoder(buffer).Encode(body)
		if nil != jsonEncoderError {
			return nil, jsonEncoderError
		}
	}
	request, requestGenerationErr := http.NewRequest(method, fullUrl.String(), buffer)
	if nil != requestGenerationErr {
		return nil, requestGenerationErr
	}
	if nil != body {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("Accept", "application/json")
	return request, nil
}

// Primmary method to make a request to the API.
func (c *Client) do(ctx context.Context, request *http.Request, v interface{}) (*http.Response, error) {
	request = request.WithContext(ctx)
	response, requestError := c.httpClient.Do(request)
	if nil != requestError {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, requestError
	}
	defer (func() { _ = response.Body.Close() })()
	parseError := json.NewDecoder(response.Body).Decode(v)
	return response, parseError

}
