# Hello World Lambda

This is a simple AWS Lambda function in Go that returns a JSON response.

## Prerequisites

- Go 1.21+
- AWS SAM CLI
- Delve (for debugging): `go install github.com/go-delve/delve/cmd/dlv@latest`

## Build

```bash
sam build --use-container
```

## Test Locally

```bash
sam local invoke HelloWorldFunction --event event.json
```

## Debug in VS Code

1. In VS Code, go to Run and Debug, select "Debug Local", and start debugging.

Set breakpoints in `HandleRequest` to debug the function. The program simulates an API Gateway request and prints the response.

## Deploy

Uncomment `lambda.Start(HandleRequest)` in `main.go` and comment out the local test code.

Then run:

```bash
sam build --use-container
sam deploy --guided
```