package lib

import (
	"github.com/logrusorgru/aurora"
	"text/template"
)

var templateFuncs template.FuncMap

func init() {
	templateFuncs = map[string]interface{}{
		"bold": aurora.Bold,
	}
}

func ParseTemplate(name, text string) *template.Template {
	t, err := template.New(name).Funcs(templateFuncs).Parse(text)
	if err != nil {
		panic(err)
	}
	return t
}
