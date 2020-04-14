package book

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/xml"
	"io"
	"os"
	"path"
	"sync"
	"text/template"
)

type Renderer struct {
	sync.RWMutex
	Reducer *minify.M
}

func NewRenderer() *Renderer {
	reducer := minify.New()
	reducer.AddFunc("text/xml", xml.Minify)
	return &Renderer{
		Reducer: reducer,
	}
}

func (r *Renderer) Render(templates []string, filename string, content interface{}, minimize bool) error {
	r.Lock()
	defer r.Unlock()
	if err := os.MkdirAll(path.Dir(filename), 0777); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	b := bytes.NewBuffer(nil)
	t := template.New(filename)
	for _, tpl := range templates {
		if _, err = t.Parse(tpl); err != nil {
			return err
		}
	}

	if err := t.Execute(b, content); err != nil {
		return err
	}
	if minimize {
		if err := r.Reducer.Minify("text/xml", file, b); err != nil {
			return err
		}
	} else {
		if _, err := io.Copy(file, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *Renderer) RenderFile(filepath, filename string, reader io.Reader) error {
	r.Lock()
	defer r.Unlock()
	if err := os.MkdirAll(path.Dir(filepath), 0777); err != nil {
		return err
	}
	file, err := os.Create(filepath + filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, reader); err != nil {
		return err
	}
	return nil
}
