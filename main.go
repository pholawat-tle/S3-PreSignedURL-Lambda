package main

import "github.com/aws/aws-lambda-go/lambda"

type Event struct {
	ItemName string `json:"item-name"`
}

type Response struct {
	PresignedURL string `json:"presigned-url"`
}

func HandleLambdaEvent(event Event) (Response, error) {
	return Response{}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
