package main

import (
	"bufio"
	"bytes"
	"github.com/c-jimin/BookMaker/core"
	"github.com/c-jimin/BookMaker/templates"
	"github.com/google/uuid"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/xml"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

func NewMobiBook() {
	b := core.NewBook("重生之2006", getCover(), readFile())
	b.Authors = []string{"雨去欲续", "CodeTech BookMaker"}
	err := b.MakeMobi()
	log.Println(err)
}

func getCover() *os.File {
	file, _ := os.Open("./templates/cover.jpg")
	return file
}

func renderFile(path, filename string, reader io.Reader) error {
	f, err := fileWriter(path + filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, reader); err != nil {
		return err
	}
	return nil
}

func fileWriter(filename string) (io.WriteCloser, error) {
	if err := os.MkdirAll(path.Dir(filename), 0777); err != nil {
		return nil, err
	}
	file, err := os.Create(filename)
	return file, err
}

func readFile() []*core.Chapter {
	chapters := make([]*core.Chapter, 0)
	file, err := os.Open("./2006.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		re, _ := regexp.Compile("<h1>(.*?)</h1>")
		title := re.FindString(txt)
		length := len(title)
		title = title[4 : length-5]
		chapter := &core.Chapter{
			Title:          title,
			FileName:       uuid.New().String(),
			Content:        strings.NewReader(txt),
			SubChapters:    nil,
			Level:          1,
			BeforeMakeFile: pretreatment,
		}
		chapters = append(chapters, chapter)
	}
	return chapters
}
func pretreatment(c *core.Chapter) error {
	content := bytes.NewBuffer(nil)
	io.Copy(content, c.Content)
	b := bytes.NewBuffer(nil)
	t := template.New(c.Title)
	if _, err := t.Parse(templates.Content); err != nil {
		return err
	}
	s := content.String()
	con := strings.Replace(s, " ", "&nbsp;&nbsp;", -1)

	if err := t.Execute(b, struct {
		Title   string
		Content string
	}{
		Title:   c.Title,
		Content: con,
	}); err != nil {
		return err
	}
	f := bytes.NewBuffer(nil)
	reducer := minify.New()
	reducer.AddFunc("text/xml", xml.Minify)
	if err := reducer.Minify("text/xml", f, b); err != nil {
		return err
	}
	c.Content = f
	return nil
}

func main() {
	NewMobiBook()
}
