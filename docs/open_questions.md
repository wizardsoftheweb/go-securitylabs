# Open Questions

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC_BY_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)

## General

* If something is marked `|Null` in the response, will it be `null` or will it be omitted when empty?

## Authentication

[Docs](https://apidocs.hunter2.com/#authentication)

* Don't the `missing`+`rotated` responses expose the API to credential stuffing? The performance of this would be [like Bogosort](https://en.wikipedia.org/wiki/Bogosort) so the risk is probably very low.
  * This assumes that I have a proper key.
  * If I have to have both a proper key and an old secret, then the risk is even lower.
  * I need valid API creds to play with this.

## Users

[Docs](https://apidocs.hunter2.com/#users-2)

* The [Get Users Details docs](https://apidocs.hunter2.com/#get-users-details) show `{current,next,previous}Url` as `/api/onboarding`; is this wrong?
* [Get User Progress](https://apidocs.hunter2.com/#get-user-progress) mentions a `page` query param. There are no examples. It also says "page of users to get" but the response format has pages of lessons.
* [Put User - Response Example](https://apidocs.hunter2.com/#put-user) misspells `disabled`.

## Summaries

[Docs](https://apidocs.hunter2.com/#summaries-2)

* [Get Progress](https://apidocs.hunter2.com/#get-progress) has an `.users[].lessons[].startTime` field that's specified as `Number`. Is it possible this should be `Number|Null`?
* [Get Progress](https://apidocs.hunter2.com/#get-progress) has an `.users[].lessons[].endTime` field that's specified as `String|Null` but is returned as, I think, a Unix timestamp.

## Roles

[Docs](https://apidocs.hunter2.com/#roles-2)

* [Get Roles](https://apidocs.hunter2.com/#get-roles) has the potential to be large. Does it paginate?
* [Get Role Progress](https://apidocs.hunter2.com/#get-role-progress) appears to return `[{}]`. However, it also has a `[].nextPage`. Since it's returning the progress for a single role, should the entire response be `[{}]` or do the pages just apply to the `users` within the output?

## Invites

[Docs](https://apidocs.hunter2.com/#invites-2)

* Can names be sent [with invites](https://apidocs.hunter2.com/#post-invites)? They don't appear in the docs.
