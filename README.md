# LarvAuth

# [![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-yellow.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/GuaiZai233/larvauth)](https://goreportcard.com/report/github.com/GuaiZai233/larvauth) [![pkg.go.dev](https://pkg.go.dev/badge/github.com/GuaiZai233/larvauth)](https://pkg.go.dev/github.com/GuaiZai233/larvauth) ![Go Version](https://img.shields.io/badge/go-1.25-blue?logo=go)

A unified account system backend, written in Go.

> This is an experimental personal learning project. Any assistance to improve are welcomed.

## Features

- **User Authentication**: Secure registration and login using JWT.
- **Database**: SQLite integration with GORM for persistent storage.
- **API**: RESTful API endpoints built with Gin framework.

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) 1.25 or higher

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/GuaiZai233/larvauth.git
    cd larvauth
    ```

2.  Install dependencies:
    ```bash
    go mod tidy
    ```

### Running the Server

#### Using Make
```bash
make build
./app
```

#### Directly with Go
```bash
go run cmd/server/main.go
```

The server will start on port `8080`.

## API Documentation

### Register
Create a new user account.

- **URL**: `/api/reg`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
    ```json
    {
        "username": "example_user",
        "password": "secure_password",
        "email": "user@example.com"
    }
    ```

### Login
Authenticate a user and receive a JWT token.

- **URL**: `/api/login`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Body**:
    ```json
    {
        "username": "example_user",
        "password": "secure_password"
    }
    ```
- **Response**:
    ```json
    {
        "message": "Login successful",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

## Contributing

Issues and PRs are welcome.

## License

[MPL 2.0](LICENSE)

---

The name of this project was inspired by this song: [larva - ガリガリさむし](https://www.youtube.com/watch?v=gKjcmeid1zQ)  (maimai 14.8 | CHUNITHM 14.7)