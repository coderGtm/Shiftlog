# release.io

A highly performant and efficient system to create, manage and publish release notes in various formats for your products. **Can easily run as a microservice on your existing server or a standalone app on a completely different system if you prefer so, no fuss invloved!!**

## Overview
release.io is a robust release note management system designed to streamline the process of creating and managing release notes for applications and products. Accessible only through user accounts, this system empowers users to efficiently handle their release documentation.

## Key Features

### 1. User Management

Access to release.io requires user authentication, ensuring secure and personalized use.

### 2. App Creation

Users have the ability to create individual applications within the system, facilitating organized release note management.
Release Versioning:

Each application can have multiple releases, allowing for meticulous version tracking and management.
### 3. Release Note Formats

release.io supports release notes in three formats: **plain text (txt), Markdown (md), and HTML (html)**. This versatile feature accommodates various documentation preferences.

![flowchart](https://github.com/coderGtm/release.io/assets/66418526/827e98d1-be10-4fdb-bff6-00b8283148bb)


## How It Works:

### 1. User Registration and Authentication:

Users register for an account, providing the necessary authentication credentials.
Creating and Managing Apps:

Once logged in, users can create and manage their applications, defining specific attributes and settings.
### 2. Versioning and Release Notes:

Within each application, users can create and manage multiple releases. For each release, they can attach release notes in their preferred format (txt, md, html).
### 3. Browsing and Exporting Release Notes:

Users can easily browse and view release notes associated with their applications. Additionally, they have the option to export notes in their chosen format.
## Benefits

- ### Efficient Documentation: 
    release.io streamlines the process of creating, updating, and managing release notes, saving valuable time for developers and product teams.

- ### Customizable Formats:
    The system accommodates different documentation styles with support for txt, md, and html formats, catering to a wide range of user preferences.

- ### Organized Versioning:
    Through intuitive version management, users can easily track the evolution of their applications and associated release notes.

- ### Enhanced Collaboration:
    The user management feature ensures that only authorized individuals have access, promoting secure collaboration within teams.

## Target Audience

Developers, product managers, and teams involved in the development and release of applications and products.
## Use Case Scenarios

#### 1. Software Development Teams:
Development teams can use release.io to maintain a clear record of version updates and communicate changes effectively.

#### 2. Product Managers:
Product managers can leverage the platform to keep stakeholders informed about releases and new features.

#### 3. Freelancers and Independent Developers:
Individual developers and freelancers can utilize release.io for efficient documentation of their projects.

***By employing release.io, users can ensure seamless release documentation, fostering a more organized and transparent development process.***

# Documentation

## Installation Instructions

### Prerequisites

- **Go (Golang):** Ensure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

- **SQLite3:** release.io requires SQLite3. If not already installed, download it from [https://www.sqlite.org/download.html](https://www.sqlite.org/download.html) and follow the installation instructions for your platform.

### Clone the Repository

Open your terminal or command prompt and navigate to your desired installation directory. Clone the repository:

```shell
git clone https://github.com/coderGtm/release.io.git
```

### Set Environment Variables

1. Locate the `demo.env` file in the project directory.
2. Open `demo.env` and fill in the required environment variables, and rename it to `.env`. This step also now protects the file via `.gitignore`.

### Download Go Modules

Navigate to the project's root directory in your terminal and run:

```shell
go mod tidy
```

### Run the Application

Start the application with the following command:

```shell
go run .
```

### Access the Application

By default, the application runs on port 8080. Access it via Postman or curl with base URL [http://localhost:8080](http://localhost:8080).

## Usage Guide

release.io offers a user-friendly interface for managing release notes. Here's a brief guide:

1. **User Registration and Login:**
   - Register for an account or log in if you already have one. User management ensures secure access.

2. **Creating Applications:**
   - After logging in, create new applications by providing details like name and description.

3. **Managing Releases:**
   - Within each application, create and manage multiple releases, organized by version numbers.

4. **Adding Release Notes:**
   - For each release, add release notes in plain text (txt), Markdown (md), or HTML (html) format.

5. **Browsing and Exporting Release Notes:**
   - Easily browse and view release notes associated with your applications.
   - Export release notes in your chosen format for documentation.

## Configuration

- Customize the application's configuration in the `main.go` file. You can change the listening port if needed.

- An empty database file named `sqlite3.db` with all necessary tables is provided, eliminating manual database setup.

- Configure environment variables in `demo.env` to match your specific requirements, and then rename it to just `.env`.