package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.S3EventRecord) (*string, error) {
	message := fmt.Sprintf("Event: %+v", event)
	log.Println(message)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
