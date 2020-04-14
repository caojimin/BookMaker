package book

import (
	"bytes"
	"github.com/c-jimin/BookMaker/errors"
	"io"
	"strconv"
)

type Chapter struct {
	Title          string
	FileName       string
	BeforeMakeFile func(*Chapter) error
	Content        io.Reader
	StaticFile     []*StaticFile
	SubChapters    []*Chapter
	Level          int // max=4
}

func (c *Chapter) MakeFile(path string) error {
	if c.Level > 4 {
		return errors.New("max level is 4 and now is", strconv.Itoa(c.Level))
	}
	if c.BeforeMakeFile != nil {
		if err := c.BeforeMakeFile(c); err != nil {
			return err
		}
	}
	if c.Content == nil {
		c.Content = bytes.NewBufferString("<body></body>")
	}
	if err := globalRenderer.RenderFile(path, c.FileName+".xhtml", c.Content); err != nil {
		return err
	}
	for _, staticFile := range c.StaticFile {
		if err := staticFile.MakeFile(path); err != nil {
			return err
		}
	}
	for _, sc := range c.SubChapters {
		if err := sc.MakeFile(path); err != nil {
			return err
		}
	}
	return nil
}
