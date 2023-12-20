package infrastructure

import (
	"bytes"
	"stori/pkg/result"
	"text/template"
)

func CreateTemplate(res result.Result) ([]byte, string) {
	tmpl, _ := template.ParseFiles("data/email.html.tmpl")
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, res)
	return tpl.Bytes(), tpl.String()
}
