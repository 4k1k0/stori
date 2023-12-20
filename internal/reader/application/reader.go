package application

import (
	"stori/internal/reader/domain"
	"stori/internal/reader/infrastructure"
)

func New(file string) domain.Reader {
	return &infrastructure.ReaderCSV{}
}
