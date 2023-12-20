package infrastructure

import (
	"stori/internal/reader/domain"
	"stori/pkg/transaction"
)

type ReaderCSV struct {
	FilePath string
	Reader   domain.Reader
}

func New(file string) *ReaderCSV {
	return &ReaderCSV{
		FilePath: file,
	}
}

func (r *ReaderCSV) Read() ([]transaction.Transaction, error) {
	return nil, nil
}
