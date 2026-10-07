# FirstMicroserviceInGo

A simple calculator microservice built using **Go** and the standard `net/http` package.

The service provides a REST API that performs basic mathematical operations such as:

* Addition
* Subtraction
* Multiplication
* Division

It also includes **Swagger/OpenAPI documentation** for testing the API.

## Project Structure

```text
FirstMicroserviceInGo/
│
├── main.go
├── go.mod
├── README.md
│
├── api/
│   ├── calculatorhandler.go
│   └── requestjson.go
│
└── docs/
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## Technologies Used

* Go
* `net/http`
* `encoding/json`
* REST API
* Swagger / OpenAPI
* JSON

## Requirements

Make sure Go is installed:

```bash
go version
```

The project currently uses:

```text
Go 1.24.5
```

## How to Run

Clone the repository and move into the project directory:

```bash
cd FirstMicroserviceInGo
```

Run the application:

```bash
go run .
```

The server will start on:

```text
http://localhost:8080
```

You should see:

```text
Server started on :8080
```

## API Endpoint

### Calculate

```text
POST /calculate
```

The API accepts two numbers and an operation.

### Request Format

```json
{
  "operation": "addition",
  "a": 5,
  "b": 10
}
```

### Response

```json
{
  "result": 15
}
```

## Supported Operations

### Addition

Request:

```json
{
  "operation": "addition",
  "a": 5,
  "b": 10
}
```

Response:

```json
{
  "result": 15
}
```

### Subtraction

Request:

```json
{
  "operation": "subtraction",
  "a": 10,
  "b": 5
}
```

Response:

```json
{
  "result": 5
}
```

### Multiplication

Request:

```json
{
  "operation": "multiplication",
  "a": 5,
  "b": 10
}
```

Response:

```json
{
  "result": 50
}
```

### Division

Request:

```json
{
  "operation": "division",
  "a": 10,
  "b": 5
}
```

Response:

```json
{
  "result": 2
}
```

## Decimal Numbers

The API also supports decimal numbers.

Example:

```json
{
  "operation": "addition",
  "a": 10.5,
  "b": 20.25
}
```

Response:

```json
{
  "result": 30.75
}
```

## Swagger Documentation

Swagger is available when the application is running.

Open:

```text
http://localhost:8080/swagger/index.html
```

Swagger provides an interactive interface where you can test the `/calculate` API without using Postman.

## Generate Swagger Documentation

If the Swagger documentation needs to be regenerated, install the Swagger CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Make sure the Go binary directory is in your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then run:

```bash
swag init
```

This generates:

```text
docs/
├── docs.go
├── swagger.json
└── swagger.yaml
```

## Error Handling

The API handles common errors such as:

### Invalid JSON

```text
Invalid JSON
```

### Invalid Number

```text
Parameter 'a' must be a number
```

or:

```text
Parameter 'b' must be a number
```

### Invalid Operation

```text
Invalid operation
```

### Division by Zero

```text
Cannot divide by zero
```

### Unsupported HTTP Method

```text
Method not allowed
```

## Example Using curl

Start the application:

```bash
go run .
```

Then execute:

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{"operation":"addition","a":5,"b":10}'
```

Expected response:

```json
{
  "result": 15
}
```

## Learning Goals

This project was created to understand the fundamentals of building a microservice in Go, including:

* Go project structure
* Packages
* Structs
* Interfaces and methods
* HTTP servers
* HTTP handlers
* REST APIs
* JSON encoding and decoding
* Request validation
* Error handling
* Type assertions
* Swagger/OpenAPI documentation
* Go modules
* Running and testing a Go microservice

## Future Improvements

Possible improvements for this project:

* Add unit tests
* Add integration tests
* Add better request validation
* Add structured logging
* Add Docker support
* Add configuration management
* Add middleware
* Add request/response models
* Add CI/CD
* Deploy the service to AWS
* Add health-check endpoint
* Add automated Swagger generation

## Author

**Ganesha Painjane**

This project is part of my journey to learn **Go backend development and microservices**.
