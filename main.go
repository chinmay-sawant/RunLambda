package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Starting Lambda function...")
	time.Sleep(5 * time.Second) // Delay to allow debugger to attach
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
	if os.Getenv("IS_LOCAL_DEBUG") == "true" {
		// Local debugging mode: read event from args and call handler directly
		if len(os.Args) > 1 {
			eventFile := os.Args[1]
			data, err := os.ReadFile(eventFile)
			if err != nil {
				fmt.Printf("Error reading event file: %v\n", err)
				return
			}
			var request events.APIGatewayProxyRequest
			if err := json.Unmarshal(data, &request); err != nil {
				fmt.Printf("Error parsing event: %v\n", err)
				return
			}
			response, err := HandleRequest(request)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Response: %+v\n", response)
		} else {
			fmt.Println("No event file provided")
		}
	} else {
		// Lambda deployment
		fmt.Println("Starting Lambda function...")
		lambda.Start(HandleRequest)
	}
}
