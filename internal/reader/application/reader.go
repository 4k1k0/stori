package application

import (
	"stori/internal/config"
	"stori/internal/reader/domain"
	"stori/internal/reader/infrastructure"
	"stori/internal/reader/infrastructure/file"
	"stori/internal/reader/infrastructure/s3"
)

func New(filepath string) domain.Reader {
	if filepath != "" {
		return &file.ReaderFile{
			File:        filepath,
			SliceReader: &infrastructure.ReaderCSV{},
		}
	}

	return &s3.S3Reader{
		Bucket:      config.Config().S3Config.Bucket,
		Key:         config.Config().S3Config.Key,
		S3Client:    config.Config().S3Config.S3Client,
		SliceReader: &infrastructure.ReaderCSV{},
	}
}
