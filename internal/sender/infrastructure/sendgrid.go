package infrastructure

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func foo(data []byte) {
	m := mail.NewV3Mail()

	from := mail.NewEmail("test", "test@example.com")
	content := mail.NewContent("text/html", "<p>Sending different attachments.</p>")
	to := mail.NewEmail("Stori", "test1@example.com")

	m.SetFrom(from)
	m.AddContent(content)

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = "Stori Test"

	m.AddPersonalizations(personalization)

	attachCSV := mail.NewAttachment()
	encoded := base64.StdEncoding.EncodeToString(data)
	attachCSV.SetContent(encoded)
	attachCSV.SetType("text/csv")
	attachCSV.SetFilename("transactions.csv")
	attachCSV.SetDisposition("attachment")

	m.AddAttachment(attachCSV)

	request := sendgrid.GetRequest(
		os.Getenv("SENDGRID_API_KEY"),
		"/v3/mail/send",
		"https://api.sendgrid.com",
	)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
