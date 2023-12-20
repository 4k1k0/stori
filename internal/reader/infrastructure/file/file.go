package file

import (
	"os"
	"stori/internal/reader/domain"
	"stori/pkg/transaction"
)

type ReaderFile struct {
	SliceReader domain.ReaderSlice
	File        string
}

func (r *ReaderFile) Read() ([]transaction.Transaction, []byte, error) {
	file, err := os.ReadFile(r.File)
	if err != nil {
		return nil, nil, err
	}

	return r.SliceReader.Read(file)
}
