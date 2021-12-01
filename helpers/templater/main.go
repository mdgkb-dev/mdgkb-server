package templater

import (
	"fmt"
	"log"
	"mdgkb/mdgkb-server/config"
	"path/filepath"
	"strings"
	"text/template"
)

type Templater struct {
	templatesPath string
}

func NewTemplater(config config.Config) *Templater {
	return &Templater{templatesPath: config.TemplatesPath}
}

func (i *Templater) Parse(templateName string, data interface{}) string {
	var buf strings.Builder
	templateName = fmt.Sprintf("%s.gohtml", filepath.Join(i.templatesPath, templateName))
	tmpl, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(&buf, data)
	strTmpl := buf.String()
	return strTmpl
}
