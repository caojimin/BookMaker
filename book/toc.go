package book

import (
	"log"
	"strconv"
)

type Toc struct {
	BookName string
	Authors  []string
	BookId   string
	MaxLevel int
	Chapters []*TocChapter
	Indent   func(int) string
}

type TocChapter struct {
	ChapterID  string
	Title      string
	FileName   string
	PlayOrder  int
	Level      int
	SubChapter []*TocChapter
	Indent     func(int) string
}

func NewToc(b *Book) *Toc {
	t := &Toc{
		BookName: b.Name,
		Authors:  b.Authors,
		BookId:   b.BookId,
	}
	maxLevel := -1
	chapters := make([]*TocChapter, 0)
	lastPlayOrder := 1
	for i, c := range b.Chapters {
		sc, lpo, level := makeTocChapters(c, "chapter", lastPlayOrder, i)
		chapters = append(chapters, sc)
		lastPlayOrder = lpo
		if level > maxLevel {
			maxLevel = level
		}
	}
	if maxLevel <= 0 || maxLevel >= 5 {
		log.Fatal("max level must be greater than 0 and less than 5, and now is ", maxLevel)
	}
	t.MaxLevel = maxLevel
	t.Chapters = chapters
	return t
}

func makeTocChapters(c *Chapter, parentChapterID string, lastPlayOrder, index int) (*TocChapter, int, int) {
	lastPlayOrder++
	t := &TocChapter{
		ChapterID: parentChapterID + "_" + strconv.Itoa(index),
		Title:     c.Title,
		FileName:  c.FileName + ".xhtml",
		PlayOrder: lastPlayOrder,
		Level:     c.Level,
	}
	maxLevel := c.Level
	subChapters := make([]*TocChapter, 0)
	for i, sc := range c.SubChapters {
		tsc, lpo, level := makeTocChapters(sc, t.ChapterID, lastPlayOrder, i)
		lastPlayOrder = lpo
		subChapters = append(subChapters, tsc)
		if level > maxLevel {
			maxLevel = level
		}
	}
	t.SubChapter = subChapters
	return t, lastPlayOrder, maxLevel
}
