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
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient(nil, nil)
	assert.NotNilf(t, client, "Client should not be nil")
	assert.Equalf(t, productionUrl, client.Config.BaseUrl, "BaseUrl should be set to productionUrl")
	clientWithConfig := NewClient(&ClientConfig{BaseUrl: &url.URL{}}, nil)
	assert.NotNilf(t, clientWithConfig, "Client should not be nil")
	assert.NotEqualf(t, productionUrl, clientWithConfig.Config.BaseUrl, "BaseUrl should not be set to productionUrl")
	clientWithHttpClient := NewClient(nil, &http.Client{Timeout: time.Duration(1) * time.Second})
	assert.NotNilf(t, clientWithHttpClient, "Client should not be nil")
	assert.NotEqualf(t, clientWithHttpClient.httpClient, http.DefaultClient, "httpClient should not be set to http.DefaultClient")
}
