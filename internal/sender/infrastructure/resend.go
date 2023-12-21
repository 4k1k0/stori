package infrastructure

import (
	"log"
	"os"

	"github.com/resendlabs/resend-go"
)

func sendEmail(emailHTML, to string) {
	apiKey := os.Getenv("STORI_SENDER_API_KEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: "Stori Test",
		Html:    emailHTML,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
