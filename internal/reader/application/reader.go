package application

import (
	"stori/internal/reader/domain"
	"stori/internal/reader/infrastructure"
)

func New(file string) domain.Reader {
	if file != "" {
		x := file.ReaderFile{
			File: file,
		}

		return &x
	}

	return &infrastructure.ReaderCSV{}
}
