![Header](release.io.png)

# release.io

A fast, efficient, and secure system to create, manage, and publish Release Notes for your apps/products.

🚀 Blazingly Fast APIs written in Go

🔖 Supports Release Notes in multiple formats *(Plain Text, Markdown and HTML)*

🔒 Fully Secure Backend with JWT Auth Tokens

![flowchart](https://github.com/coderGtm/release.io/assets/66418526/827e98d1-be10-4fdb-bff6-00b8283148bb)

## 🛠️ Getting Started

release.io is available as an API system that can be integrated into an existing server on a different port or deployed as a stand-alone microservice. A custom frontend is needed to access the Dashboard as of now.

#### Prerequisites

- **Go (Golang):** Ensure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

- **SQLite3:** release.io requires SQLite3. If not already installed, download it from [https://www.sqlite.org/download.html](https://www.sqlite.org/download.html) and follow the installation instructions for your platform.

#### Clone the Repository

Open your terminal or command prompt and navigate to your desired installation directory. Clone the repository:

```shell
git clone https://github.com/coderGtm/release.io.git
```

#### Set Environment Variables

1. Locate the `demo.env` file in the project directory.
2. Open `demo.env` and fill in the required environment variables.
3. Rename `demo.env` to `.env`. This step now protects the file via `.gitignore`.

#### Download Go Modules

Navigate to the project's root directory in your terminal and run:

```shell
go mod tidy
```

#### Run the Application

Start the application with the following command:

```shell
go run .
```

#### Access the Application

By default, the application (API) runs on port 8080. You can access it via Postman or curl with the base URL [http://localhost:8080](http://localhost:8080). A very bare-bones and rather "incomplete" frontend is provided in the [frontend](frontend) directory which is routed to the base url scheme, i.e. you can access it via [http://localhost:8080/signup](http://localhost:8080/signup) scheme. It is **highly reccomended and maybe necessary** that you develop your custom frontend, while taking reference from what is provided. _As a side note, it would be great if you can update the frontend in this repo with your improved version._


## 🚀 Deployment

You can change the port on which the service runs by modifying the `PORT` value in your environment config file (`.env`).

## 🚧 How It Works:

### 1. User Registration and Authentication:

Users register for an account, providing the necessary authentication credentials.

### 2. Creating and Managing Apps:

Once logged in, users can create and manage their applications, defining specific attributes and settings.

### 3. Versioning and Release Notes:

Within each application, users can create and manage multiple releases. For each release, they can attach release notes in their preferred format (text, MD, HTML).

### 4. Browsing and Exporting Release Notes:

Users can easily browse and view release notes associated with their applications. Additionally, they have the option to export notes in their chosen format.

## ✅ Benefits

- ### Efficient Documentation: 
    release.io streamlines the process of creating, updating, and managing release notes, saving valuable time for developers and product teams.

- ### Customizable Formats:
    The system accommodates different documentation styles with support for text, md, and HTML formats, catering to a wide range of user preferences.

- ### Organized Versioning:
    Through intuitive version management, users can easily track the evolution of their applications and associated release notes.

- ### Enhanced Collaboration:
    The user management feature ensures that only authorized individuals have access, promoting secure collaboration within teams.

## 📝 API Documentation

***Please see [documentation.md](documentation.md) for a detailed API Documentation of release.io***

## 🔍 Built With

* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go. It features a martini-like API with performance that is up to 40 times faster
* [jwt-go](https://github.com/golang-jwt/jwt/v5) - A go implementation of JSON Web Tokens.
* [go-sqlite3](https://github.com/mattn/go-sqlite3) - A SQLite3 driver that conforms to the built-in database/SQL interface.
* [bluemonday](https://github.com/microcosm-cc/bluemonday) - bluemonday is an HTML sanitizer implemented in Go. It is fast and highly configurable.
* [Go Cryptography](https://golang.org/x/crypto) - This repository holds supplementary Go cryptography libraries.

## 🌱 Contributing

Contributions are welcome. Please open an issue first and refer to that in your pull request.

## 🌟 Support

You can show your support by becoming a ***stargazer...*** 

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
