package log

import (
	"os"
	"testing"
	"text/template"
)

type TestStruct struct {
	Name string
	Age  uint
}

func TestRegister(t *testing.T) {

	TestStructTemplate := template.Must(template.New("TestStruct").Parse(`
####
  Name {{ .Name }} ({{ .Age }})
####
`))

	TS := TestStruct{}
	Register(&TS, TestStructTemplate)
	Register(TS, TestStructTemplate)

	if len(templates) != 1 {
		t.Log("Pointer mishandled. There should be only on template.")
		t.Log(templates)
		t.Fail()
	}

	if templates["github.com/omeid/logskin.TestStruct"] == nil {
		t.Log("Template failed to register.")
		t.Log(templates)
		t.Fail()
	}

	IntTemplate := template.Must(template.New("TestStruct").Parse(` ### This is an int {{ . }} ###
	
	`))
	Register(int(1), IntTemplate)

}

func TestParse(t *testing.T) {
	l := New(os.Stdout, "", 0)

	l.Skin(TestStruct{"Hello", 21})
	l.Skin(5)
}
