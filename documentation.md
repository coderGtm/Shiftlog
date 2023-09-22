## This Page documents the APIs provided by release.io

release.io features a set of APIs to easily create, manage, and publish Release Notes for your apps/products. The APIs are divided into 4 major parts:

## 1. Auth


* ### `/api/createAccount`

    #### Used to create a new account.

    **TYPE:** GET

    **Headers:**

    _NONE_

    **Parameters:**

    - `username: string`

    - `password: string`

    **Response:**

    - **BadRequest (400):** _Empty parameters in Request Body_

    - **BadRequest (400):** _Illegal username or password!_

    - **Conflict (409):** _Username Already Exists_

    - **OK (200):** _`{username: "<username>", authToken: "<authToken>"}`_



* ### `/api/deleteAccount`

    #### Used to delete an existing account.

    **TYPE:** DELETE

    **Headers:**

    - Authorization Bearer : `authToken`

    **Parameters:**

    _NONE_

    **Response:**

    - ****Unauthorized (401):**** _Auth token missing!_

    - ****Unauthorized (401):**** _Invalid Auth Token_

    - **OK (200):** _User Account Deleted Successfully!_



* ### `/api/updateUsername`

    #### Used to update the username of an account.

    **TYPE:** PUT

    **Headers:**

    - Authorization Bearer : `authToken`

    **Parameters:**

    - `newUsername: string`

    **Response:**

    - ****Unauthorized (401):**** _Auth token missing!_

    - ****Unauthorized (401):**** _Invalid Auth Token_

    - **BadRequest (400):** _Empty parameters in Request Body_

    - **BadRequest (400):** _Illegal username provided!_

    - **Conflict (409):** _This username is already taken!_

    - **OK (200):** _Username updated successfully!_



* ### `/api/updatePassword`

    #### Used to update the password of an account.

    **TYPE:** PUT

    **Headers:**

    - Authorization Bearer : `authToken`

    **Parameters:**

    `newPassword: string`

    **Response:**

    - ****Unauthorized (401):**** _Auth token missing!_

    - ****Unauthorized (401):**** _Invalid Auth Token_

    - **BadRequest (400):** _Empty parameters in Request Body_

    - **BadRequest (400):** _Password contain illegal charachters!_

    - **Conflict (409):** _New Password cannot be same as old Password!_

    - **OK (200):** _`{authToken: <authToken>}`_



* ### `/api/login`

    #### Used to login to an existing account.

    **TYPE:** POST

    **Headers:**

    _NONE_

    **Parameters:**

    `username: string`

    `password: string`

    **Response:**

    - **BadRequest (400):** _Empty parameters in Request Body_

    - **BadRequest (400):** _Illegal username or password!_

    - **Conflict (409):** _Username does not exist!_

    - **Unauthorized (401):** _Invalid login credentials!_

    - **OK (200):** _`{authToken: <authToken>}`_



* ### `/api/logout`

    #### Used to logout an account.

    **TYPE:** GET

    **Headers:**

    - Authorization Bearer : `authToken`

    **Parameters:**

    _NONE_

    **Response:**

    - ****Unauthorized (401):**** _Auth token missing!_

    - ****Unauthorized (401):**** _Invalid Auth Token_

    - **OK (200):** _Logged out!_

    

## 2. Dashboard
## 3. App
## 4. Release