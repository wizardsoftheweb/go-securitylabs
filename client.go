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
	"net/http"
	"net/url"
)

// Get the latest prod URL: https://apidocs.hunter2.com/#production
var productionUrl, _ = url.Parse("https://securitylabs.veracode.com/api/")

type ClientConfig struct {
	BaseUrl *url.URL
}

type Client struct {
	Config     *ClientConfig
	httpClient *http.Client
}

func NewClient(config *ClientConfig, httpClient *http.Client) *Client {
	var newConfig *ClientConfig
	var newHttpClient *http.Client
	if nil == config {
		newConfig = &ClientConfig{
			BaseUrl: productionUrl,
		}
	} else {
		newConfig = config
	}
	if nil == httpClient {
		newHttpClient = http.DefaultClient
	} else {
		newHttpClient = httpClient
	}
	return &Client{
		Config:     newConfig,
		httpClient: newHttpClient,
	}
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
