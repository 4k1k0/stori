package s3

import (
	"context"
	"io"
	"log"
	"stori/internal/reader/domain"
	"stori/pkg/transaction"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Reader struct {
	Bucket      string
	Key         string
	S3Client    *s3.Client
	SliceReader domain.ReaderSlice
}

func (s *S3Reader) Read() ([]transaction.Transaction, []byte, error) {
	log.Println("Read from s3 reader")
	objOut, err := s.S3Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &s.Bucket,
		Key:    &s.Key,
	})

	if err != nil {
		log.Println("there was an error with getobject")
		return nil, nil, err
	}

	defer objOut.Body.Close()

	fileContent, err := io.ReadAll(objOut.Body)
	if err != nil {
		log.Println("there was an error with readAll function")
		return nil, nil, err
	}

	return s.SliceReader.Read(fileContent)
}
