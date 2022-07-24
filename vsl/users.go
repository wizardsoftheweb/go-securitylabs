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
)

const (
	GetUsersPath        = "/users"
	GetUsersDetailsPath = "/users/details"
	GetUserProgressPath = "/users/%s/progress"
)

// GetUsersUsersRoles
// This is a list of roles attached to a user
// https://apidocs.hunter2.com/#get-users
type GetUsersUsersRoles []string

// GetUsersUsers
// These are all the properties of a user from the /api/user endpoint
// https://apidocs.hunter2.com/#get-users
type GetUsersUsers struct {
	Id         string             `json:"id"`
	IsAdmin    bool               `json:"isAdmin"`
	IsDisabled bool               `json:"isDisabled"`
	Email      string             `json:"email"`
	Joined     bool               `json:"joined"`
	LastActive int64              `json:"lastActive"`
	Roles      GetUsersUsersRoles `json:"roles"`
}

// GetUsersResponse
// This is the full body of the /api/user endpoint
// https://apidocs.hunter2.com/#get-users
type GetUsersResponse struct {
	NextPage *string         `json:"nextPage"`
	Users    []GetUsersUsers `json:"users"`
}

// GetUsersOptions
// These are the query params for the /api/users endpoint
// https://apidocs.hunter2.com/#get-users
type GetUsersOptions struct {
	Page *int `query:"page"`
}

func (c *Client) GetUsers(ctx context.Context, options *GetUsersOptions) ([]GetUsersUsers, error) {
	// The only way to generate an error from Client.newRequest is if the body can't build
	// Since we have no body, we can safely ignore the error
	request, _ := c.newRequest(http.MethodGet, GetUsersPath, options, nil)
	var responseBody GetUsersResponse
	// TODO: Verify error makes sense once Client.do has been fully tested
	_, responseError := c.do(ctx, request, &responseBody)
	if nil != responseError {
		return nil, responseError
	}
	return responseBody.Users, nil
}

// GetUsersDetailsPages
// This is the format of the pages object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type GetUsersDetailsPages struct {
	Current     int     `json:"current"`
	Previous    *int    `json:"previous"`
	Next        *int    `json:"next"`
	Limit       int     `json:"limit"`
	Total       int     `json:"total"`
	CurrentUrl  string  `json:"currentUrl"`
	NextUrl     *string `json:"nextUrl"`
	PreviousUrl *string `json:"previousUrl"`
}

// GetUsersDetailsUsersRoles
// This is the format of the roles object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type GetUsersDetailsUsersRoles struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetUsersDetailsUsers
// This is the format of the users object on the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type GetUsersDetailsUsers struct {
	Id                       string                      `json:"id"`
	LastActive               *int64                      `json:"lastActive"`
	LabsCompleted            int                         `json:"labsCompleted"`
	PercentRequiredCompleted float64                     `json:"percentRequiredCompleted"`
	Points                   int                         `json:"points"`
	Name                     string                      `json:"name"`
	Roles                    []GetUsersDetailsUsersRoles `json:"roles"`
}

// GetUsersDetailsResponse
// This is the full body of the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type GetUsersDetailsResponse struct {
	Pages GetUsersDetailsPages   `json:"pages"`
	Users []GetUsersDetailsUsers `json:"users"`
}

// GetUsersDetailsOptions
// These are the query params for the /api/users/details endpoint
// https://apidocs.hunter2.com/#get-users-details
type GetUsersDetailsOptions struct {
	CampaignIds []string `query:"campaignIds"`
	EndTime     *int64   `query:"endTime"`
	Limit       *int     `query:"limit"`
	Page        *int     `query:"page"`
	Phrase      *string  `query:"phrase"`
	StartTime   *int64   `query:"startTime"`
	RoleIds     []string `query:"roleIds"`
	Sort        *string  `query:"sort"`
	SortType    *string  `query:"sortType"`
}

func (c *Client) GetUsersDetails(ctx context.Context, options *GetUsersDetailsOptions) (GetUsersDetailsResponse, error) {
	// The only way to generate an error from Client.newRequest is if the body can't build
	// Since we have no body, we can safely ignore the error
	request, _ := c.newRequest(http.MethodGet, GetUsersDetailsPath, options, nil)
	var responseBody GetUsersDetailsResponse
	// TODO: Verify error makes sense once Client.do has been fully tested
	_, responseError := c.do(ctx, request, &responseBody)
	if nil != responseError {
		return GetUsersDetailsResponse{}, responseError
	}
	return responseBody, nil
}

// GetUserProgressLesson
// This is the format of the lessons object on the /api/users/:id/progress endpoint
// https://apidocs.hunter2.com/#get-user-progress
type GetUserProgressLesson struct {
	Module      string  `json:"module"`
	LessonId    string  `json:"lessonId"`
	LessonName  string  `json:"lessonName"`
	LastVisited string  `json:"lastVisited"`
	Status      string  `json:"status"`
	Minutes     float64 `json:"minutes"`
	Points      int     `json:"points"`
	StartRating int     `json:"startRating"`
	EndRating   *int    `json:"endRating"`
}

// GetUserProgressResponse
// This is the full body of the /api/users/:id/progress endpoint
// https://apidocs.hunter2.com/#get-user-progress
type GetUserProgressResponse struct {
	PointsRequired int                     `json:"pointsRequired"`
	PointsPossible int                     `json:"pointsPossible"`
	Lessons        []GetUserProgressLesson `json:"lessons"`
}

// TODO: The API docs call out pages as an optional parameter; how does this work?
//type GetUserProgressOptions struct {
//	Page *int `query:"page"`
//}

func (c *Client) GetUserProgress(ctx context.Context, userId string) (GetUserProgressResponse, error) {
	// The only way to generate an error from Client.newRequest is if the body can't build
	// Since we have no body, we can safely ignore the error
	request, _ := c.newRequest(http.MethodGet, fmt.Sprintf(GetUserProgressPath, userId), nil, nil)
	var responseBody GetUserProgressResponse
	// TODO: Verify error makes sense once Client.do has been fully tested
	_, responseError := c.do(ctx, request, &responseBody)
	if nil != responseError {
		return GetUserProgressResponse{}, responseError
	}
	return responseBody, nil
}

// PutUser
// This is the format of both the request and repsonse body of the
// /api/user/:id endpoint
// https://apidocs.hunter2.com/#put-user
type PutUser struct {
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	Admin    bool     `json:"admin"`
	Disabled bool     `json:"disabled"`
	RoleIds  []string `json:"roleIds"`
}
