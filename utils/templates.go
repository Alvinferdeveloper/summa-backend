package utils

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// ParseTemplate parses an HTML template with the given data.
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(filepath.Join("templates", templateFileName))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
