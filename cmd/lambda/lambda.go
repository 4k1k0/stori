package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Filename string `json:"filename,omitempty"`
	Email    string `json:"email,omitempty"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Filename %q! Email %q", event.Filename, event.Email)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
