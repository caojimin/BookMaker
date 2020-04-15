package jsonbook

import (
	"bytes"
	"encoding/json"
	"github.com/c-jimin/BookMaker/book"
	"github.com/google/uuid"
	"io"
)

type JsonBook struct {
	BookName string         `json:"bookname"`
	Authors  []string       `json:"authors"`
	Chapters []*JsonChapter `json:"chapters"`
}

type JsonChapter struct {
	Title string   `json:"title"`
	Lines []string `json:"lines"`
}

func (jb *JsonBook) NewBook() *book.Book {
	return &book.Book{
		Name:       jb.BookName,
		Authors:    jb.Authors,
		BookId:     uuid.New().String(),
		Cover:      nil,
		Chapters:   jb.genChapters(),
		TempPath:   "./" + uuid.New().String() + "/",
		OutputPath: "./output/",
		Gen:        book.DefaultGen,
		Renderer:   book.NewRenderer(),
		PreChecker: book.DefaultPreChecker.Check,
	}
}

func (jb *JsonBook) genChapters() []*book.Chapter {
	chapters := make([]*book.Chapter, 0)
	for _, c := range jb.Chapters {
		content := bytes.NewBuffer(nil)
		content.Write([]byte(Strings("<h2>", c.Title, "</h2>")))
		for _, line := range c.Lines {
			content.Write([]byte(Strings("<p>", line, "</p>")))
		}
		chapters = append(chapters, &book.Chapter{
			Title:          c.Title,
			FileName:       uuid.New().String(),
			BeforeMakeFile: book.DefaultPreprocessor.Do,
			Content:        content,
			StaticFile:     nil,
			SubChapters:    nil,
			Level:          1,
		})
	}
	return chapters
}

func New(r io.Reader) *JsonBook {
	b := bytes.NewBuffer(nil)
	io.Copy(b, r)
	jb := new(JsonBook)
	if err := json.Unmarshal(b.Bytes(), jb); err != nil {
		panic(err)
	}
	return jb
}
