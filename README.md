# User Management API - Tester Challenge

This project provides a simple User Management API implemented in Go using the Gin framework. The API allows you to register users, list all users, and delete users.

## Prerequisites

- Go programming language (version 1.16 or later)
- Git (optional if you prefer to download the code manually)

## Getting Started

1. Clone the repository using the following command:

   ```bash
   git clone <repository_url>
   ```

   Alternatively, you can manually download the project code from the repository.

2. Navigate to the project directory:

   ```bash
   cd go-application
   ```

3. Build and run the API:

   ```bash
   go run .
   ```

   The API will start running at `http://localhost:8080`.

4. Test the API endpoints:

   - To register a user, send a POST request to `http://localhost:8080/register` with a JSON payload containing the user details, e.g., `{"id": "1", "name": "John Doe", "age": 25}`.

   - To list all users, send a GET request to `http://localhost:8080/users`.

   - To delete a user, send a DELETE request to `http://localhost:8080/users/:id`, where `:id` is the ID of the user you want to delete.

## Project Structure

The project structure is as follows:

- `go-applicaton`: Project root directory.
    - `main.go`: Main entry point of the application containing the API endpoints and server setup.
    - `users.json`: JSON file used to store user data persistently.
- `README.md`: Project documentation.

## Dependencies

The project uses the following dependencies:

- Gin: Web framework for building the API.