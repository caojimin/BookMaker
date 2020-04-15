package book

import (
	"bytes"
	"github.com/c-jimin/BookMaker/templates"
	"github.com/google/uuid"
	"io"
	"strings"
	"text/template"
)

type Preprocessor struct {
	pretreatments []func(*Chapter, string) (string, error)
}

func NewPreprocessor(funcs ...func(*Chapter, string) (string, error)) *Preprocessor {
	return &Preprocessor{
		funcs,
	}
}

func (p *Preprocessor) Do(c *Chapter) error {
	b := bytes.NewBuffer(nil)
	io.Copy(b, c.Content)
	str := b.String()
	var err error
	for _, f := range p.pretreatments {
		if str, err = f(c, str); err != nil {
			return err
		}
	}
	return nil
}

func ReplaceSpace(_ *Chapter, s string) (string, error) {
	return strings.Replace(s, " ", "&nbsp;&nbsp;", -1), nil
}

func Format2xhtml(c *Chapter, s string) (string, error) {
	b := bytes.NewBuffer(nil)
	t := template.New(uuid.New().String())
	if _, err := t.Parse(templates.Content); err != nil {
		return "", err
	}
	if err := t.Execute(b, struct {
		Title   string
		Content string
	}{
		Title:   c.Title,
		Content: s,
	}); err != nil {
		return "", err
	}
	return b.String(), nil
}

func Minimize(_ *Chapter, s string) (string, error) {
	return globalRenderer.Reducer.String("text/xml", s)
}

func SetContent(c *Chapter, s string) (string, error) {
	c.Content = strings.NewReader(s)
	return s, nil
}

var DefaultPreprocessor = NewPreprocessor(ReplaceSpace, Format2xhtml, Minimize, SetContent)
