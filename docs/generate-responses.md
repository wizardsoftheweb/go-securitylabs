# Generating Responses

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC_BY_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Authentication](#authentication)
  - [403 Invalid](#403-invalid)
  - [404 Missing](#404-missing)
- [Everything Else](#everything-else)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Authentication

- [Docs](https://apidocs.hunter2.com/#authentication)
- [Code](../authentication_test.go)

### 403 Invalid

```bash
$ curl -s https://securitylabs.veracode.com/api/users | jq
{
  "message": "ApiCredential invalid"
}
```

### 404 Missing

The creds are from [the docs](https://apidocs.hunter2.com/#authentication).

```bash
$ curl -s https://securitylabs.veracode.com/api/users -H "auth: 486a4f5d52daafac2aff4891156d9993400716e2:1e07a337896d4cbe60d1d143" | jq
{
  "message": "ApiCredential missing"
}
```

## Everything Else

If it's not on this page, I haven't been able to generate the request. See [my notes](../README.md#notes).
