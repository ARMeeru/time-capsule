# Time Capsule

Welcome to the **Time Capsule** project! This application allows users to create messages (capsules) that will be delivered to them at a specified future date and time. The project features both an API and a minimal web interface for user interaction.

---

## Table of Contents

- [Time Capsule](#time-capsule)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Running the Project Locally](#running-the-project-locally)
    - [Running the Main Version (API Only)](#running-the-main-version-api-only)
    - [Running the Web Version](#running-the-web-version)
    - [Switching Between Main and Web Version](#switching-between-main-and-web-version)
  - [API Documentation](#api-documentation)
    - [Authentication Endpoints](#authentication-endpoints)
      - [Register a New User](#register-a-new-user)
      - [Login](#login)
    - [Capsule Endpoints](#capsule-endpoints)
      - [Create a New Capsule](#create-a-new-capsule)
      - [Get All Capsules (Future Feature)](#get-all-capsules-future-feature)
  - [Project Structure](#project-structure)
  - [License](#license)

---

## Project Overview

The Time Capsule project is a Go-based application that allows users to:

- Register and log in to their accounts.
- Create time capsules containing messages to be delivered at a specified future date.
- Receive capsules via email when the delivery time arrives.

The application provides both a RESTful API and a minimal web interface built with the Gin framework.

---

## Features

- **User Authentication**: Secure registration and login functionality with password hashing.
- **Capsule Creation**: Users can create messages to be delivered at a future date.
- **Scheduler**: A background scheduler checks for capsules to be delivered and sends them via email.
- **Web Interface**: A minimal web version for user interaction through a browser.
- **API Endpoints**: RESTful API for integration with other applications or front-end clients.

---

## Technologies Used

- **Go (Golang)**: The primary programming language.
- **Gin**: Web framework used for handling HTTP requests.
- **GORM**: ORM library for database interactions.
- **SQLite**: Database for storing user and capsule data.
- **JWT**: JSON Web Tokens for authentication.
- **Bcrypt**: Password hashing algorithm.
- **Go Modules**: Dependency management.

---

## Getting Started

### Prerequisites

- **Go**: Version 1.16 or higher is required. [Download Go](https://golang.org/dl/)
- **Git**: For version control. [Download Git](https://git-scm.com/downloads)
- **SQLite**: The application uses SQLite as the database. [Download SQLite](https://www.sqlite.org/download.html) (if not already installed)

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/time-capsule.git
   cd time-capsule
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the project root directory with the following content:

   ```env
   APP_PORT=8080
   JWT_SECRET=your_jwt_secret_key
   EMAIL_HOST=smtp.example.com
   EMAIL_PORT=587
   EMAIL_USERNAME=your_email@example.com
   EMAIL_PASSWORD=your_email_password
   ```

   - Replace `your_jwt_secret_key` with a secure secret key for JWT.
   - Replace email configurations with your SMTP server details.

3. **Install Dependencies**

   ```bash
   go mod tidy
   ```

---

## Running the Project Locally

### Running the Main Version (API Only)

The main version provides API endpoints for user registration, login, and capsule creation.

1. **Switch to the Main Branch**

   ```bash
   git checkout main
   ```

2. **Run the Application**

   ```bash
   go run main.go
   ```

3. **Test the API**

   Use tools like **Postman** or **cURL** to interact with the API endpoints. See the [API Documentation](#api-documentation) section for endpoint details.

### Running the Web Version

The web version includes a minimal web interface for user interaction via a browser.

1. **Switch to the Web Version Branch**

   ```bash
   git checkout web-version
   ```

2. **Run the Application**

   ```bash
   go run main.go
   ```

3. **Access the Application**

   - Open your browser and navigate to `http://localhost:8080/register` to create a new account.
   - After registering, log in at `http://localhost:8080/login`.
   - Create a capsule at `http://localhost:8080/capsules/new`.

### Switching Between Main and Web Version

To switch between the main (API-only) and web versions of the application, use Git branches:

- **Switch to Main Version (API Only)**

  ```bash
  git checkout main
  ```

- **Switch to Web Version**

  ```bash
  git checkout web-version
  ```

---

## API Documentation

### Authentication Endpoints

#### Register a New User

- **URL**: `/register`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Request Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  - `201 Created` on success.
  - `400 Bad Request` if the email is already registered or data is invalid.

#### Login

- **URL**: `/login`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Request Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  - `200 OK` with a JSON containing the JWT token.
  - `401 Unauthorized` if credentials are invalid.

### Capsule Endpoints

#### Create a New Capsule

- **URL**: `/capsules`
- **Method**: `POST`
- **Headers**:

  - `Authorization: Bearer <JWT_TOKEN>`
  - `Content-Type: application/json`

- **Request Body**:

  ```json
  {
    "message": "Your message here",
    "deliver_at": "2024-12-31T23:59:00"
  }
  ```

  - `deliver_at` should be in ISO 8601 format.

- **Response**:

  - `201 Created` on success.
  - `400 Bad Request` if data is invalid.
  - `401 Unauthorized` if JWT token is missing or invalid.

#### Get All Capsules (Future Feature)

- **URL**: `/capsules`
- **Method**: `GET`
- **Headers**:

  - `Authorization: Bearer <JWT_TOKEN>`

- **Response**:

  - `200 OK` with a JSON array of capsules.
  - `401 Unauthorized` if JWT token is missing or invalid.

---

## Project Structure

```
time-capsule/
├── config/
│   └── config.go          // Configuration loading
├── controllers/
│   ├── auth_controller.go // Authentication handlers
│   └── capsule_controller.go // Capsule handlers
├── middlewares/
│   └── auth_middleware.go // JWT authentication middleware
├── models/
│   ├── user.go            // User model
│   └── capsule.go         // Capsule model
├── routes/
│   └── routes.go          // Route definitions
├── utils/
│   ├── db.go              // Database connection
│   ├── email.go           // Email sending utility
│   ├── scheduler.go       // Scheduler for capsule delivery
│   └── template.go        // Template rendering helper
├── templates/
│   ├── base.html          // Base HTML template
│   ├── register.html      // Registration page template
│   ├── login.html         // Login page template
│   └── create_capsule.html // Capsule creation page template
├── .env                   // Environment variables (not committed)
├── .gitignore             // Git ignore file
├── go.mod                 // Go module file
├── go.sum                 // Go dependencies lock file
└── main.go                // Application entry point
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---