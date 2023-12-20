package file

import (
	"log"
	"os"
	"stori/internal/reader/domain"
	"stori/pkg/transaction"
)

type ReaderFile struct {
	SliceReader domain.ReaderSlice
	File        string
}

func (r *ReaderFile) Read() ([]transaction.Transaction, error) {
	log.Println("read method from ReaderFile")

	file, err := os.ReadFile(r.File)
	if err != nil {
		return nil, err
	}

	return r.SliceReader.Read(file)
}
