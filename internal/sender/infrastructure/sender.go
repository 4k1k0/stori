package infrastructure

import (
	"log"
	"stori/pkg/result"
)

type EmailSender struct {
	Email string
}

func (e *EmailSender) Send(r result.Result, data []byte) error {
	log.Println("Send email from EmailSender")

	_, _ = CreateTemplate(r)
	return nil
}
