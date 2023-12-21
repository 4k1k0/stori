package infrastructure

import (
	"bytes"
	"stori/internal/config"
	"stori/pkg/result"
	"text/template"
)

func CreateTemplate(res result.Result) string {
	cfg := config.Config()
	tmpl, _ := template.ParseFS(cfg.FS, "assets/email.html.tmpl")

	var tpl bytes.Buffer
	tmpl.Execute(&tpl, res)
	return tpl.String()
}
