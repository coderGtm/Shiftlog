## 📋 This Page documents the APIs provided by release.io

release.io features a set of APIs to easily create, manage, and publish Release Notes for your apps/products. The APIs are divided into 4 major parts:

## 1. 🔐 Auth

As the name suggests, these endpoints are used for all Authentication related stuff.

* ### 📝 `/api/createAccount`

    #### Used to create a new account.

    **TYPE:** GET

    **Headers:**

    - 🔒 _NONE_

    **Parameters:**

    - `username: string`

    - `password: string`

    **Response:**

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Illegal username or password!_

    - ⚠️ **Conflict (409):** _Username Already Exists_

    - ✅ **OK (200):** _`{username: "<username>", authToken: "<authToken>"}`_



* ### 🗑️ `/api/deleteAccount`

    #### Used to delete an existing account.

    **TYPE:** DELETE

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - 🔒 _NONE_

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ✅ **OK (200):** _User Account Deleted Successfully!_



* ### ✏️ `/api/updateUsername`

    #### Used to update the username of an account.

    **TYPE:** PUT

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `newUsername: string`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Illegal username provided!_

    - ⚠️ **Conflict (409):** _This username is already taken!_

    - ✅ **OK (200):** _Username updated successfully!_



* ### 🔑 `/api/updatePassword`

    #### Used to update the password of an account.

    **TYPE:** PUT

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    `newPassword: string`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Password contain illegal charachters!_

    - ⚠️ **Conflict (409):** _New Password cannot be same as old Password!_

    - ✅ **OK (200):** _`{authToken: <authToken>}`_



* ### 🔑 `/api/login`

    #### Used to login to an existing account.

    **TYPE:** POST

    **Headers:**

    - 🔒 _NONE_

    **Parameters:**

    `username: string`

    `password: string`

    **Response:**

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Illegal username or password!_

    - ⚠️ **Conflict (409):** _Username does not exist!_

    - 🚫 **Unauthorized (401):** _Invalid login credentials!_

    - ✅ **OK (200):** _`{username: "<username>", authToken: <authToken>}`_



* ### 🚪 `/api/logout`

    #### Used to logout an account.

    **TYPE:** GET

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - 🔒 _NONE_

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ✅ **OK (200):** _Logged out!_



## 2. 📊 Dashboard

These endpoints are typically to be consumed by the Dashboard of a frontend, hence they are classified as Dashboard APIs.

* ### 📥 `/api/getApps`

    #### Used to get list of all apps of a user.

    **TYPE:** GET

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - 🔒 _NONE_

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ✅ **OK (200):** _`[{id:<id>, name: <name>, hidden: [true/false], createdAt: <timestamp>, updatedAt: <timestamp>}, ...]`_



* ### ➕ `/api/createApp`

    #### Used to get list of all apps of a user.

    **TYPE:** POST

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `appName: string`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Empty app Name is not allowed._

    - ❌ **BadRequest (400):** _Illegal app Name_

    - ✅ **OK (200):** _`{id:<id>, name: <name>, hidden: [true/false], createdAt: <timestamp>, updatedAt: <timestamp>}`_



* ### 🗑️ `/api/deleteApp`

    #### Used to delete an app of a user.

    **TYPE:** DELETE

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `appId: int`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _appId must be an Integer._

    - ❌ **BadRequest (400):** _Illegal app Id_

    - 🚫 **Unauthorized (401):** _Unauthorized deletion!_

    - ✅ **OK (200):** _`App deleted successfully!`_



* ### ✏️ `/api/updateApp`

    #### Used to delete an app of a user.

    **TYPE:** PUT

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `appId: int`
    - `name: string`
    - `hidden: int`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Illegal values provided!_

    - ❌ **BadRequest (400):** _Hiddden parameter must have a 'true' or 'false' value_

    - ❌ **BadRequest (400):** _appId must be an Integer._

    - 🚫 **Unauthorized (401):** _Unauthorized update!_

    - ✅ **OK (200):** _`App updated successfully!`_



## 3. 📱 App

These are App-specific APIs, typically used to present an "App page" by a frontend, hence the classification.

* ### 🔄 `/api/getRelease`

    #### Used to get Releases of an App.

    **TYPE:** GET

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `appId: int`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Illegal app Id_

    - ❌ **BadRequest (400):** _App ID must be an Integer._

    - 🚫 **Unauthorized (401):** _Unauthorized access!_

    - ✅ **OK (200):** _`[{id:<id>, appId: <appId>, versionCode: <code>, versionName: <name>, notesTxt: <txt>, notesMd: <md>, notesHtml: <html>, data: <stringData>, hidden: [true/false], createdAt: <timestamp>, updatedAt: <timestamp>}, ...]`_



* ### ➕ `/api/createReleases`

    #### Used to create a new Release of an App.

    **TYPE:** POST

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `appId: int`
    - `versionName: string`
    - `versionCode: int`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Illegal input parameter values_

    - ❌ **BadRequest (400):** _App ID must be an integer_

    - ❌ **BadRequest (400):** _Version Code must be an integer_

    - ❌ **BadRequest (400):** _Empty Version Name is not allowed._

    - ❌ **BadRequest (400):** _This Version Code already exists_

    - 🚫 **Unauthorized (401):** _Unauthorized Request!_

    - ✅ **OK (200):** _`{id:<id>, appId: <appId>, versionCode: <code>, versionName: <name>, notesTxt: <txt>, notesMd: <md>, notesHtml: <html>, data: <stringData>, hidden: [true/false], createdAt: <timestamp>, updatedAt: <timestamp>}`_



* ### 🗑️ `/api/deleteRelease`

    #### Used to delete an existing Release of an App.

    **TYPE:** DELETE

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `releaseId: int`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Illegal Release Id_

    - ❌ **BadRequest (400):** _releaseId must be an Integer._

    - 🚫 **Unauthorized (401):** _Delete Request Unauthorized!_

    - ✅ **OK (200):** _`Release deleted successfully!`_



* ### ✏️ `/api/updateRelease`

    #### Used to update details of a Release.

    **TYPE:** PUT

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `releaseId: int`
    - `versionName: string`
    - `versionCode: int`
    - `hidden: int`
    - `data: string`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Empty parameters in Request Body_

    - ❌ **BadRequest (400):** _Illegal values provided!_

    - ❌ **BadRequest (400):** _Hiddden parameter must have a 'true' or 'false' value_

    - ❌ **BadRequest (400):** _Release Id must be an Integer._

    - ❌ **BadRequest (400):** _Version Code must be an Integer._

    - 🚫 **Unauthorized (401):** _Unauthorized update!_

    - ✅ **OK (200):** _`Release Details updated successfully!`_



## 4. 🚀 Release

These endpoints are Release-specific, in the manner that they are concerned about managing and publishing of a specific Release. 

* ### 📖 `/api/getReleaseNotes`    _[Unprotected Endpoint]_

    #### Used to get Release Notes for a Release.

    **NOTE:** This endpointis an unprotected endpoint, i.e, no authentication is required to access it.

    This API supports fetching Release Notes via 2 methods, ordered by precedence:

    1. Directly by Release ID
    2. By App ID and Version Code _(Note that you can also use the "latest" keyword in `versionCode` to automatically select the Release with the highest Version Code for given `appId`)_

    **TYPE:** GET

    **Headers:**

    - 🔒 _NONE_

    **Parameters:**

    - `releaseId: int`
    - `appId: int`
    - `versionCode: int`
    - `hidden: int`
    - `data: string`

    **Response:**

    - ❌ **BadRequest (400):** _Missing Parameters!_

    - ❌ **BadRequest (400):** _Illegal Release ID_

    - ❌ **BadRequest (400):** _Release ID must be an Integer!_

    - ❌ **BadRequest (400):** _Illegal App ID or Version Code_

    - ❌ **BadRequest (400):** _App ID must be an Integer!_

    - ❌ **BadRequest (400):** _Invalid Version Code_

    - **NotFound (404):** _Release Notes not found!_

    - ✅ **OK (200):** _`{versionCode: <code>, versionName: <name>, notesTxt: <txt>, notesMd: <md>, notesHtml: <html>, updatedAt: <timestamp>}`_



* ### ✏️ `/api/updateReleaseNotes`

    #### Used to update Release Notes of a Release.

    **TYPE:** PUT

    **Headers:**

    - 🔑 Authorization Bearer : `authToken`

    **Parameters:**

    - `releaseId: int`
    - `notesTxt: string`
    - `notesMd: string`
    - `notesHtml: string`

    **Response:**

    - 🚫 **Unauthorized (401):** _Auth token missing!_

    - 🚫 **Unauthorized (401):** _Invalid Auth Token_

    - ❌ **BadRequest (400):** _Missing Release ID_

    - ❌ **BadRequest (400):** _Illegal values for Release ID provided!_

    - ❌ **BadRequest (400):** _Release Id must be an Integer._

    - 🚫 **Unauthorized (401):** _Unauthorized update!_

    - ✅ **OK (200):** _`Release Notes updated successfully!`_