package utils

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"html/template"
	"path/filepath"
	"reflect"
)

//ParseTemplate -> to parse the template with given data
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("utils/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	unescapedString := html.UnescapeString(body)
	return unescapedString, nil
}

//IsInterfaceEmpty -> to check if the interface is empty
func IsInterfaceEmpty(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
