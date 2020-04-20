package book

import (
	"github.com/c-jimin/BookMaker/errors"
	"github.com/c-jimin/BookMaker/templates"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

type Book struct {
	Name     string
	Authors  []string
	BookId   string
	Cover    io.ReadCloser
	Chapters []*Chapter

	// 临时文件路径，在make完成后会被cleanup
	TempPath string

	// OutputPath 和 OutputWriter 二选其一, 如果 OutputPath 设置了则会优先使用, 否则将会将文件写入到 OutputWriter。
	// 如果 OutputPath 和 OutputWriter 都没有设置将会返回错误
	OutputPath   string
	OutputWriter io.WriteCloser

	// 生成器
	Gen func(*Book) error

	// 渲染器 这个东西准备改掉 => type interface{}
	Renderer *Renderer
}

func New(name string, chapters []*Chapter) *Book {
	return &Book{
		Name:     name,
		Authors:  []string{"CodeTech BookMaker"},
		Chapters: chapters,
	}
}

// 预处理 用于检查book的各项参数，给部分参数赋初值
func (b *Book) pretreatment() error {
	// 先查错
	if b.Name == "" {
		return errors.New("book name is empty")
	}
	if b.OutputPath == "" && b.OutputWriter == nil {
		return errors.New("both outputPath and outputWriter are empty")
	}

	// 再赋值
	flag := true
	for _, author := range b.Authors {
		if author == "CodeTech BookMaker" {
			flag = false
			break
		}
	}
	if flag {
		b.Authors = append(b.Authors, "CodeTech BookMaker")
	}
	if b.BookId == "" {
		b.BookId = uuid.New().String()
	}
	if b.Cover == nil {
		file, err := os.Open("./static/cover.jpg")
		if err != nil {
			return err
		}
		b.Cover = file
	}
	if b.TempPath == "" {
		b.TempPath = "./" + uuid.New().String() + "/"
	}
	var err error
	b.TempPath, err = filepath.Abs(b.TempPath)
	if err != nil {
		return err
	}
	if b.OutputPath != "" {
		b.OutputPath, err = filepath.Abs(b.OutputPath)
		if err != nil {
			return err
		}
	}

	if b.Gen == nil {
		b.Gen = DefaultGen
	}
	if b.Renderer == nil {
		b.Renderer = globalRenderer
	}
	return nil
}

func (b *Book) MakeEPub() error {
	defer b.cleanup()
	if err := b.pretreatment(); err != nil {
		return err
	}
	return errors.New("make ePub already not implemented yet")
}

func (b *Book) MakeMobi() error {
	defer b.cleanup()
	if err := b.pretreatment(); err != nil {
		return err
	}
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
	return b.output(filename)
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
		filepath.Join(b.TempPath, "toc.xhtml"),
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
		filepath.Join(b.TempPath, b.Name+".opf"),
		toc, true,
	)
}

func (b *Book) makeNcx(toc *Toc) error {
	return b.Renderer.Render(
		[]string{
			templates.TocXml,
			templates.NavPoint,
		},
		filepath.Join(b.TempPath, "toc.ncx"),
		toc, true,
	)
}

func (b *Book) makeFile() error {
	defer b.Cover.Close()
	if err := b.Renderer.RenderFile(filepath.Join(b.TempPath, "cover.jpg"), b.Cover); err != nil {
		return err
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

func (b *Book) output(filename string) error {
	if b.OutputPath != "" {
		// 直接修改path将文件"剪切"过去
		if err := os.MkdirAll(b.OutputPath, 0755); err != nil {
			return err
		}
		tf := filepath.Join(b.TempPath, filename)
		of := filepath.Join(b.OutputPath, filename)
		if err := os.Rename(tf, of); err != nil {
			return err
		}
	} else if b.OutputWriter != nil {
		defer b.OutputWriter.Close()
		tf := filepath.Join(b.TempPath, filename)
		file, err := os.Open(tf)
		if err != nil {
			return err
		}
		_, err = io.Copy(b.OutputWriter, file)
		return err
	} else {
		return errors.New("both outputPath and outputWriter are empty")
	}
	return nil
}

func (b *Book) cleanup() {
	_ = os.RemoveAll(b.TempPath)
}
