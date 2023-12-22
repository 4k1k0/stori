package infrastructure

import (
	"log"
	"stori/internal/config"
	"stori/pkg/result"
)

type EmailSender struct {
	Email string
}

func (e *EmailSender) Send(r result.Result, data []byte) error {
	log.Println("Send email from EmailSender")
	html := CreateTemplate(r)
	sendEmail(html, config.Config().Email)
	return nil
}
