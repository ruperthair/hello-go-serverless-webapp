package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	head = map[string]string{"Content-Type": "text/plain"}
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s1 := fmt.Sprintf("Method = %s\nBody size = %d\nHeaders:\n", request.HTTPMethod, len(request.Body))
	for key, value := range request.Headers {
		s1 += fmt.Sprintf("  %s: %s\n", key, value)
	}
	return events.APIGatewayProxyResponse{Body: s1, StatusCode: 200, Headers: head}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
