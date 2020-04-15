package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/c-jimin/BookMaker/book"
	"github.com/c-jimin/BookMaker/book/jsonbook"
	"github.com/c-jimin/BookMaker/templates"
	"github.com/google/uuid"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/xml"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func NewMobiBook() {
	b := book.New("重生之2006", getCover(), readFile())
	b.Authors = []string{"雨去欲续", "CodeTech BookMaker"}
	err := b.MakeMobi()
	log.Println(err)
}

func getCover() *os.File {
	file, _ := os.Open("./templates/cover.jpg")
	return file
}

func readFile() []*book.Chapter {
	chapters := make([]*book.Chapter, 0)
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
		chapter := &book.Chapter{
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
func pretreatment(c *book.Chapter) error {
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

func bookFromJson() {
	file, err := os.Open("./2006.json")
	if err != nil {
		panic(err)
	}
	jb := jsonbook.New(file)
	fmt.Println(jb.NewBook().MakeMobi())
}

func main() {
	bookFromJson()
}
