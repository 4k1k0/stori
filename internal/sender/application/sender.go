package application

import (
	"stori/internal/sender/domain"
	"stori/internal/sender/infrastructure"
)

func New(email string) domain.Sender {
	return &infrastructure.EmailSender{
		Email: email,
	}
}
