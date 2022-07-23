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
	"strings"
)

const (
	testKey           = "proper"
	testSecret        = "secret"
	testRotatedSecret = "rotated"
)

// Authorization responses pulled from
// https://apidocs.hunter2.com/#authentication
// I have no idea if these are actually what the API returns
func authenticationHandlerForTests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authKey, secret string
		for key, value := range r.Header {
			if strings.ToLower(key) == "auth" {
				exploded := strings.Split(value[0], ":")
				authKey, secret = exploded[0], exploded[1]
				break
			}
		}
		if "" == authKey || "" == secret {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if testKey != authKey {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("{\"message\":\"missing\"}"))
			return
		}
		if testSecret != secret && testRotatedSecret != secret {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("{\"message\":\"disabled\"}"))
			return
		}
		if testRotatedSecret == secret {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("{\"rotated\":\"disabled\"}"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
