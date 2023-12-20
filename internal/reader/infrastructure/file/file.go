package file

import (
	"log"
	"os"
	"stori/pkg/transaction"
)

type ReaderFile struct {
	File string
}

func (r *ReaderFile) Read() ([]transaction.Transaction, error) {
	log.Printf("Read method frmo file struct | %q", r.File)
	file, err := os.ReadFile(r.File)

	if err != nil {
		return nil, err
	}

	log.Println(file)

	return nil, nil
}
