package book

import (
	"github.com/c-jimin/BookMaker/errors"
	"github.com/c-jimin/BookMaker/templates"
	"github.com/google/uuid"
	"io"
	"os"
)

type Book struct {
	Name       string
	Authors    []string
	BookId     string
	Cover      io.ReadCloser
	Chapters   []*Chapter
	TempPath   string
	OutputPath string
	Gen        func(*Book) error
	Renderer   *Renderer
}

func New(name string, cover io.ReadCloser, chapters []*Chapter) *Book {
	return &Book{
		Name:       name,
		Authors:    []string{"CodeTech BookMaker"},
		BookId:     uuid.New().String(),
		Cover:      cover,
		Chapters:   chapters,
		TempPath:   "./" + uuid.New().String() + "/",
		OutputPath: "./output/",
		Gen:        DefaultGen,
		Renderer:   globalRenderer,
	}
}

func (b *Book) MakeEPub() error {
	defer b.cleanup()
	return errors.New("make ePub already not implemented yet")
}

func (b *Book) MakeMobi() error {
	defer b.cleanup()
	if err := b.makeToc(); err != nil {
		return err
	}
	if err := b.makeFile(); err != nil {
		return err
	}
	filename := b.Name + ".mobi"
	if err := b.Gen(b); err != nil {
		// 有可能含有警告，返回码不为0，判断mobi文件是否生成
		if _, fileErr := os.Stat(b.TempPath + filename); fileErr != nil {
			return err
		}
	}
	_ = os.Mkdir(b.OutputPath, 0755)
	if err := os.Rename(b.TempPath+filename, b.OutputPath+filename); err != nil {
		return err
	}
	return nil
}

func (b *Book) makeToc() error {
	toc := NewToc(b)
	if err := b.makeNcx(toc); err != nil {
		return err
	}
	if err := b.makeOpf(toc); err != nil {
		return err
	}
	return b.Renderer.Render(
		[]string{
			templates.TocXhtml,
			templates.TOCChapter,
		},
		b.TempPath+"toc.xhtml",
		toc, true,
	)
}

func (b *Book) makeOpf(toc *Toc) error {
	return b.Renderer.Render(
		[]string{
			templates.OpfXml,
			templates.ManifestItem,
			templates.SpineItemRef,
		},
		b.TempPath+b.Name+".opf",
		toc, true,
	)
}

func (b *Book) makeNcx(toc *Toc) error {
	return b.Renderer.Render(
		[]string{
			templates.TocXml,
			templates.NavPoint,
		},
		b.TempPath+"toc.ncx",
		toc, true,
	)
}

func (b *Book) makeFile() error {
	if b.Cover != nil {
		defer b.Cover.Close()
		if err := b.Renderer.RenderFile(b.TempPath, "cover.jpg", b.Cover); err != nil {
			return err
		}
	}
	if b.Chapters != nil {
		for _, chapter := range b.Chapters {
			if err := chapter.MakeFile(b.TempPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *Book) cleanup() {
	_ = os.RemoveAll(b.TempPath)
}
