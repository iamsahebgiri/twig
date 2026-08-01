package renderer

import (
	"html/template"
	"strings"
)

func funcMap() template.FuncMap {
	return template.FuncMap{
		"upper": strings.ToUpper,
	}
}
