package templates

import (
	"github.com/aymerick/raymond"
	"io/ioutil"
	"io"
)

var parsedTemplates map[string]*raymond.Template

func Init() error {

	parsedTemplates = make(map[string]*raymond.Template)

	var err error
	parsedTemplates["account.history"], err = loadTemplate("./templates/grilla.template.html")
	if err != nil {
		return err
	}
	return nil
}

func loadTemplate(file string) (*raymond.Template, error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tpl, err := raymond.Parse(string(data))
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

func Stream(templateID string, ctx interface{}, writer io.Writer) error {
	out, err := parsedTemplates[templateID].Exec(ctx)
	if err != nil { return err }
	_, err = writer.Write([]byte(out))
	return nil
}