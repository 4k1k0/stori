package infrastructure

import (
	"log"
	"stori/pkg/result"
)

type EmailSender struct {
	Email string
}

func (e *EmailSender) Send(r result.Result) error {
	log.Println("Send email from EmailSender")
	return nil
}

func (e *EmailSender) format(r result.Result) []byte {
	return nil
}
