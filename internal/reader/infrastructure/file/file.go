package file

import (
	"log"
	"os"
	"stori/pkg/transaction"
)

type ReaderFile struct {
	Path string
}

func (r *ReaderFile) Read() ([]transaction.Transaction, error) {
	file, err := os.ReadFile(r.Path)

	if err != nil {
		return nil, err
	}

	log.Println(file)

	return nil, nil
}
