package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"message":"Hello World"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	// Local debugging simulation
	request := events.APIGatewayProxyRequest{
		Resource:   "/hello",
		Path:       "/hello",
		HTTPMethod: "GET",
		Headers:    map[string]string{"Content-Type": "application/json"},
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "test",
		},
	}
	response, err := HandleRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %+v\n", response)

	// Uncomment for Lambda deployment
	// lambda.Start(HandleRequest)
}
