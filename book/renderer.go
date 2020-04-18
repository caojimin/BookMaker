package book

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/xml"
	"io"
	"os"
	"path/filepath"
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
	fp, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	dir, _ := filepath.Split(fp)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	b := bytes.NewBuffer(nil)
	t := template.New(fp)
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

func (r *Renderer) RenderFile(path string, reader io.Reader) error {
	r.Lock()
	defer r.Unlock()
	fp, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	dir, _ := filepath.Split(fp)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, reader); err != nil {
		return err
	}
	return nil
}
