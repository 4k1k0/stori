package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func HandleRequest(ctx context.Context, event events.S3Event) (*string, error) {
	message := fmt.Sprintf("Event: %+v", event)
	log.Println(os.Environ())
	log.Println(message)

	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("failed to load default config: %s", err)
		return &message, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.URLDecodedKey

	objOut, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		log.Println("could not load file with get object function")
		return &message, err
	}

	defer objOut.Body.Close()

	x, err := io.ReadAll(objOut.Body)
	if err != nil {
		log.Println("could not read the file with readAll function")
		return &message, err
	}

	fileString := string(x)

	log.Println("file:")
	log.Println(fileString)

	headOutput, err := s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Printf("error getting head of object %s/%s: %s", bucket, key, err)
		return &message, err
	}
	log.Printf("successfully retrieved %s/%s of type %s", bucket, key, *headOutput.ContentType)

	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
