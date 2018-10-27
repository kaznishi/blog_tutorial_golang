package controller

import (
	"html/template"
	"strings"
)

func initializeTemplate() *template.Template{
	TemplateFunctions := map[string]interface{}{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}
	return template.New("dummy").Funcs(TemplateFunctions)
}
