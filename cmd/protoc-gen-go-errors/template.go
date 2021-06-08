package main

import (
	"bytes"
	"text/template"
)

var errorsTemplate = `
{{ range .Errors }}

func Is{{.Value}}(err error) bool {
	e := errors.FromError(err)
	if e.Reason == {{.Name}}_{{.Value}}.String() && e.Code == {{.HttpCode}} {
		return true
	}
	return false
}

func Error{{.Value}}(format string, args ...interface{}) *errors.Error {
	 return errors.New({{.HttpCode}}, {{.Name}}_{{.Value}}.String(), fmt.Sprintf(format, args...))
}

{{- end }}
`

type errorInfo struct {
	Name  string
	Value string
	HttpCode int
}
type errorWrapper struct {
	Errors []*errorInfo
}

func (e *errorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}