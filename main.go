package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	HTTPMethodNotSupported = errors.New("no name was provided in the HTTP body")
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s1 := fmt.Sprintf("Body size = %d. \n", len(request.Body))
	s1 += fmt.Sprintf("Headers:\n")
	for key, value := range request.Headers {
		s1 += fmt.Sprintf("  %s: %s\n", key, value)
	}
	if request.HTTPMethod == "GET" {
		s1 += fmt.Sprintf("\nGET METHOD\n")
		return events.APIGatewayProxyResponse{Body: s1, StatusCode: 200}, nil
	} else if request.HTTPMethod == "POST" {
		s1 += fmt.Sprintf("\nPOST METHOD\n")
		return events.APIGatewayProxyResponse{Body: s1, StatusCode: 200}, nil
	} else {
		fmt.Printf("NEITHER\n")
		return events.APIGatewayProxyResponse{}, HTTPMethodNotSupported
	}
}

func main() {
	lambda.Start(HandleRequest)
}
