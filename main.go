package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Event struct {
	BucketName string `json:"bucket-name"`
	ItemName   string `json:"item-name"`
}

type Response struct {
	PresignedURL string `json:"presigned-url"`
}

func HandleLambdaEvent(event Event) (Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		return Response{}, err
	}

	params := s3.GetObjectInput{
		Bucket: &event.BucketName,
		Key:    &event.ItemName,
	}
	presignedClient := s3.NewPresignClient(s3.NewFromConfig(cfg), s3.WithPresignExpires(1*time.Minute))

	object, err := presignedClient.PresignGetObject(context.TODO(), &params)
	if err != nil {
		return Response{}, err
	}

	return Response{PresignedURL: object.URL}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
