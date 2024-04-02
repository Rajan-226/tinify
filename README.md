# Tinify README

This README provides information about the endpoints available in the Go application.

## Endpoints

### Tinify Endpoint

- **Method:** POST
- **Path:** /v1/tinify
- **Handler Function:** `controllers.Tinify`
- **Description:** This endpoint is used to perform URL shortening.

### Metrics Endpoint

- **Method:** GET
- **Path:** /v1/metrics
- **Handler Function:** `controllers.Metrics`
- **Description:** This endpoint retrieves metrics related to URL shortening.

### Redirect Endpoint

- **Method:** GET
- **Path:** /v1/{path}
- **Handler Function:** `controllers.Redirect`
- **Description:** This endpoint redirects shortened URLs to their original long URLs.

## Usage

To use the endpoints, make HTTP requests to the corresponding URLs using the specified methods.

### Example Usage

#### Tinify Endpoint

```bash
curl -X POST http://localhost:8080/v1/tinify -d '{"url": "https://example.com"}'
```

Metrics Endpoint
```bash
curl http://localhost:8080/v1/metrics
```
Redirect Endpoint
```bash
curl http://localhost:8080/v1/abc123
```
Replace http://localhost:8080 with the actual base URL of your Go application

## Running the Server
To run the server, execute the following command in your terminal:

```bash
go run cmd/api/main.go
```
This command will start the Go server, allowing you to access the defined endpoints.

If you don't have Redis installed on Docker Desktop, you can follow these steps to set it up:
```bash
docker pull redis
docker run --name my-redis -d redis
```

