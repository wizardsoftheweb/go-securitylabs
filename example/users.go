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

package example

import (
	"context"
	"fmt"

	"github.com/wizardsoftheweb/go-securitylabs/vsl"
)

func getUsers() {
	client := vsl.NewClient(nil, nil)
	client.SetAuthFromEnvironment()
	users1, err1 := client.GetUsers(context.Background(), nil)
	if nil != err1 {
		panic(err1)
	}
	for _, user := range users1 {
		fmt.Printf("%+v\n", user)
	}
	page := new(int)
	*page = 1
	users2, err2 := client.GetUsers(context.Background(), &vsl.GetUsersOptions{Page: page})
	if nil != err2 {
		panic(err2)
	}
	for _, user := range users2 {
		fmt.Printf("%+v\n", user)
	}
}
