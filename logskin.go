package log

import (
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"text/template"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
}

func name(value interface{}) string {
	rt := reflect.TypeOf(value)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	name := ""
	if rt.PkgPath() != "" {
		name += rt.PkgPath()
	}

	if rt.Name() != "" {
		name += "." + rt.Name()
	}

	return name
}

func Register(value interface{}, Template *template.Template) {

	//Does this looks good?
	if err := Template.Execute(ioutil.Discard, value); err != nil {
		panic(err)
	}

	templates[name(value)] = Template
}

type Logger struct {
	*log.Logger
	out io.Writer
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{log.New(out, prefix, flag), out}
}

func (l *Logger) Skin(value interface{}) bool {

	t, ok := templates[name(value)]
	if !ok {
		return false
	}

	if err := t.Execute(l.out, value); err != nil {
		panic(err)
	}

	return true
}
