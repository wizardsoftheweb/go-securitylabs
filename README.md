# Veracode Security Labs Go Client

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC_BY_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Overview](#overview)
- [References](#references)
- [Notes](#notes)
- [Tentative Roadmap](#tentative-roadmap)
  - [Library](#library)
  - [Housekeeping](#housekeeping)
    - [`golangci-lint`](#golangci-lint)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

My goal for this package is to provide a simple Go client for the Veracode Security Labs API.

```go
import "github.com/wizardsoftheweb/go-securitylabs/vsl"
```

Check out [the examples](./example) for usage.

## References

- [Veracode Security Labs API Docs](https://apidocs.hunter2.com/)
- [Writing a Go client for your RESTful API](https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c)
- [Adding context and options to your Go client package](https://medium.com/@marcus.olsson/adding-context-and-options-to-your-go-client-package-244c4ad1231b)

## Notes

I don't have access to a Veracode Security Labs account for testing. My current employer wanted me to write a proposal to be able to develop against our account off-hours. That's work and I don't like doing work for hobby code. If you're interested in sponsoring me or providing access to an account I can run tests against, feel free to reach out!

Sometime soon I'll have all the GitHub niceties like a Contributing.md and issues templates.

I have [several open questions](docs/open_questions.md).

## Tentative Roadmap

None of these are in any particular order.

### Library

- [x] Get something simple pulled out of the wrapper article
- [x] Learn how to use `httptest.Server`
- [x] Build request and response structs for [each of the available URLs](https://apidocs.hunter2.com/#endpoints) (where applicable)
  - [x] Users
    - [x] GET /api/users?page=0
    - [x] GET /api/users/details?page=0
    - [x] GET /api/users/:id/progress
    - [x] PUT /api/users/:id
  - [x] Summaries
    - [x] GET /api/onboarding?page=0
    - [x] GET /api/progress?page=0
    - [x] GET /api/campaigns/progress?page=0
    - [x] GET /api/engagement/time
  - [x] Lessons
    - [x] GET /api/lessons?page=0
    - [x] GET /api/lessons/:id/progress?page=0
    - [x] GET /api/lessons/search
  - [x] Roles
    - [x] GET /api/roles
    - [x] GET /api/roles/:id/progress?page=0
  - [x] Invites
    - [x] POST /api/invites
- [ ] Collect all (as many as possible) of the responses from input permutations on [the available URLs](https://apidocs.hunter2.com/#endpoints) (needs a real account)
  - [ ] Authentication
  - [ ] Users
    - [ ] GET /api/users?page=0
    - [ ] GET /api/users/details?page=0
    - [ ] GET /api/users/:id/progress
    - [ ] PUT /api/users/:id
    - [ ] DELETE /api/users/:id
  - [ ] Summaries
    - [ ] GET /api/onboarding?page=0
    - [ ] GET /api/progress?page=0
    - [ ] GET /api/campaigns/progress?page=0
    - [ ] GET /api/engagement/time
  - [ ] Lessons
    - [ ] GET /api/lessons?page=0
    - [ ] GET /api/lessons/:id/progress?page=0
    - [ ] GET /api/lessons/search
  - [ ] Roles
    - [ ] GET /api/roles
    - [ ] GET /api/roles/:id/progress?page=0
  - [ ] Invites
    - [ ] POST /api/invites
- [ ] Mock all [the available URLs](https://apidocs.hunter2.com/#endpoints) (these still need to be tested against a real account to verify branches; anything marked off has only been mocked against the docs)
  - [x] Authentication
  - [x] Users
    - [x] GET /api/users?page=0
    - [x] GET /api/users/details?page=0
    - [x] GET /api/users/:id/progress
    - [x] PUT /api/users/:id
    - [x] DELETE /api/users/:id
  - [ ] Summaries
    - [ ] GET /api/onboarding?page=0
    - [ ] GET /api/progress?page=0
    - [ ] GET /api/campaigns/progress?page=0
    - [ ] GET /api/engagement/time
  - [ ] Lessons
    - [ ] GET /api/lessons?page=0
    - [ ] GET /api/lessons/:id/progress?page=0
    - [ ] GET /api/lessons/search
  - [ ] Roles
    - [ ] GET /api/roles
    - [ ] GET /api/roles/:id/progress?page=0
  - [ ] Invites
    - [ ] POST /api/invites
- [ ] Develop wrappers for [each endpoint](https://apidocs.hunter2.com/#endpoints) (including bubbling up errors)
  - [ ] Authentication
  - [ ] Users
    - [ ] GET /api/users?page=0
    - [ ] GET /api/users/details?page=0
    - [ ] GET /api/users/:id/progress
    - [ ] PUT /api/users/:id
    - [ ] DELETE /api/users/:id
  - [ ] Summaries
    - [ ] GET /api/onboarding?page=0
    - [ ] GET /api/progress?page=0
    - [ ] GET /api/campaigns/progress?page=0
    - [ ] GET /api/engagement/time
  - [ ] Lessons
    - [ ] GET /api/lessons?page=0
    - [ ] GET /api/lessons/:id/progress?page=0
    - [ ] GET /api/lessons/search
  - [ ] Roles
    - [ ] GET /api/roles
    - [ ] GET /api/roles/:id/progress?page=0
  - [ ] Invites
    - [ ] POST /api/invites
- [ ] Develop convenience methods for the API calls (eg `GetAllUsersDetails`)

### Housekeeping

- [ ] Set up CI pipelines (GHA? CircleCI?)
- [ ] Define nice status checks like code coverage
- [ ] Figure out godoc

#### `golangci-lint`

- [ ] Reenable
  - [ ] `unused`
  - [ ] `deadcode`
- [ ] Follow `structcheck` issue fix for Go 1.18
