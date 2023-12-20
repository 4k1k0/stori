package application

import (
	"stori/internal/reader/domain"
	"stori/internal/reader/infrastructure"
	"stori/internal/reader/infrastructure/file"
)

func New(filepath string) domain.Reader {
	if filepath != "" {
		return &file.ReaderFile{
			File:        filepath,
			SliceReader: &infrastructure.ReaderCSV{},
		}
	}

	return nil
}
