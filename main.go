package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Inside the handler")
	fmt.Printf("Request method: %s\n", request.HTTPMethod)
	fmt.Printf("Request path: %s\n", request.Path)
	fmt.Printf("Request body: %s\n", request.Body)
	fmt.Println("Processing request...")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"message":"Hello World"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {

	// Uncomment for Lambda deployment
	fmt.Println("Starting Lambda function...")
	lambda.Start(HandleRequest)
	fmt.Println("Lambda function started.")
}
