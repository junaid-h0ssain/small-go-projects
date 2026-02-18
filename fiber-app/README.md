# fiber-app

A Go project that provides a basic CRM-like API for managing leads using the Fiber web framework and GORM with SQLite.

## Features
-   **Fiber framework** for fast and efficient HTTP routing.
-   **GORM** for Object-Relational Mapping (ORM).
-   **SQLite** as the database database for storage.

## Installation and Run

1.  Navigate to the `fiber-app` folder:
    ```bash
    cd fiber-app
    ```
2.  Install dependencies:
    ```bash
    go mod tidy
    ```
3.  Run the application:
    ```bash
    go run main.go
    ```
4.  The server should start on port `8080`.

## API Endpoints

-   `GET /`: List all leads.
-   `GET /:id`: Retrieve details for a lead with a specific ID.
-   `POST /`: Create a new lead.
-   `DELETE /:id`: Delete a lead with a specific ID.

## Data Model (Lead)

-   `Name` (string)
-   `Description` (string)
-   `Company` (string)
-   `Email` (string)
-   `Phone` (string)
