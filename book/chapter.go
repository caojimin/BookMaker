package book

import (
	"bytes"
	"github.com/c-jimin/BookMaker/errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

type Chapter struct {
	Title    string
	FileName string

	// 预处理器
	// 在makefile之前调用，用于修改Chapter内的任何内容，一般用于对章节内容格式修改或者排版或者去除广告等
	BeforeMakeFile func(*Chapter) error
	Content        io.ReadCloser
	StaticFile     []*StaticFile
	SubChapters    []*Chapter
	Level          int // max=4
}

func (c *Chapter) MakeFile(path string) error {
	if c.BeforeMakeFile != nil {
		if err := c.BeforeMakeFile(c); err != nil {
			return err
		}
	}
	if c.Level > 4 {
		return errors.New("max level is 4 and now is", strconv.Itoa(c.Level))
	}
	if c.Content == nil {
		c.Content = ioutil.NopCloser(bytes.NewBufferString("<body></body>"))
	}
	defer c.Content.Close()
	if err := globalRenderer.RenderFile(filepath.Join(path, c.FileName+".xhtml"), c.Content); err != nil {
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
