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

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

My goal for this package is to provide a simple Go client for the Veracode Security Labs API.

## References

- [Veracode Security Labs API Docs](https://apidocs.hunter2.com/)
- [Writing a Go client for your RESTful API](https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c)

## Notes

I don't have access to a Veracode Security Labs account for testing. My current employer wanted me to write a proposal to be able to develop against our account off-hours. That's work and I don't like doing work for hobby code. If you're interested in sponsoring me or providing access to an account I can run tests against, feel free to reach out!

Sometime soon I'll have all the GitHub niceties like a Contributing.md and issues templates.

## Tentative Roadmap

None of these are in any particular order.

### Library

- [ ] Get something simple pulled out of the wrapper article
- [ ] Learn how to use `httptest.Server`
- [ ] Mock all the available URLs
- [ ] Develop wrappers for each endpoint (expand this list?)

### Housekeeping

- [ ] Set up CI pipelines (GHA? CircleCI?)
- [ ] Define nice status checks like code coverage
