package view

import "html/template"

var GenerateFuns = template.FuncMap{
	"part":  part,
	"url":   url,
	"yield": url,
}

func part(file string) string {
	return ""
}

func url(file string) string {
	return ""
}
