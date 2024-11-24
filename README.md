# FP PBKK Golang

This repository contains the final project for the Framework-Based Programming course, developed using the Go programming language.

## Authors

|    NRP     |         Nama         |
| :--------: | :------------------: |
| 5025221002 | Iftala Zahri Sukmana |
| 5025221252 |     Revy Pramana     |

## Quick Start

To set up and run the project locally, follow these steps:

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Revprm/go-fp-pbkk.git
   ```

2. **Navigate to the Project Directory:**

   ```bash
   cd go-fp-pbkk
   ```

3. **Copy the Example Environment File:**

   ```bash
   cp .env.example .env
   ```

4. **Install Dependencies:**

   Ensure you have [Go](https://golang.org/dl/) installed. Then, run:

   ```bash
   go mod tidy
   ```

5. **Run Database Migrations:**

   ```bash
   go run main.go --migrate
   ```

6. **Start the Application:**

    ```bash
    go run main.go
    ````

The application should now be running on `http://localhost:8080`.

## Program Flow

The application follows a clean architecture pattern, ensuring separation of concerns and maintainability.

`request` > router > controller > service > repository > service > controller > router > `response`

## Commands

- **Run the Application:**

  ```bash
  go run main.go
  ```

- **Run Database Migrations:**

  ```bash
  go run main.go --migrate
  ```

- **Run Tests:**

  ```bash
  go run main.go --test
  ```

## Directory Structure

The project is organized as follows:

```
go-fp-pbkk/
├── command/            # CLI commands
├── config/             # Configuration files
├── constants/          # Constant values used across the application
├── controller/         # HTTP request handlers (controller layer)
├── dto/                # Data Transfer Objects
├── entity/             # Database entities (data layer)
├── helpers/            # Utility functions
├── middleware/         # HTTP middleware
├── migrations/         # Database migration files
├── repository/         # Data access layer
├── routes/             # API route definitions
├── script/             # Auxiliary scripts
├── service/            # Business logic layer (service layer)
├── tests/              # Test cases
├── utils/              # Utility functions
└── main.go             # Entry point of the application
```

## Response Format

```json
{
  "status": true | false,
  "message": string,
  "error": null | "Error description (for failures)",
  "data": null | "Payload of the response (optional)",
  "meta": null | "Pagination metadata (optional)"
}
```

**Response Structure:**

1. Success
2. Error
3. Pagination

Success

```json
{
    "status": true,
    "message": string,
    "data": [] or {} (optional)
}

```

Error

```json
{
  "status": false,
  "message": string,
  "error": string,
  "data": null
}

```

Pagination

```json
{
  "status": true,
  "success": string,
  "data": PaginationData{}
}
```

PaginationData

```json
{
  "page": number,
  "per_page": number,
  "max_page": number,
  "count": number
}
```
