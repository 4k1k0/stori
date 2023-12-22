package main

import (
	"context"
	"embed"
	"log"
	"os"

	storiConfig "stori/internal/config"
	database "stori/internal/database/application"
	processor "stori/pkg/processor/application"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:embed all:assets/*
var assets embed.FS

func getS3Config(ctx context.Context) (*s3.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Println("could not load default config")
		return nil, err
	}
	return s3.NewFromConfig(sdkConfig), nil
}

func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {

	s3Client, err := getS3Config(ctx)
	if err != nil {
		return "", err
	}

	db, err := database.New().Connect()
	if err != nil {
		log.Println("there was an error with the database")
		return "", nil
	}

	defer storiConfig.Config().Database.Close()

	storiConfig.New(storiConfig.Cfg{
		FS:       assets,
		Email:    os.Getenv("STORI_EMAIL"),
		Database: db,
		S3Config: storiConfig.S3Config{
			Bucket:   event.Records[0].S3.Bucket.Name,
			Key:      event.Records[0].S3.Object.URLDecodedKey,
			S3Client: s3Client,
		}})

	_, err = processor.New().Process()
	if err != nil {
		log.Println("there was an error during the execution")
		return "", err
	}

	return "ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}
