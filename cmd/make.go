package main

import (
	"github.com/c-jimin/BookMaker/book"
	"github.com/c-jimin/BookMaker/book/jsonbook"
	"github.com/c-jimin/BookMaker/errors"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
)

var cmdMake = &cli.Command{
	Name:  "make",
	Usage: "make book from a file",
	Action: func(c *cli.Context) error {
		coverFile, err := os.Open(c.String("cover"))
		if err != nil {
			return err
		}
		_type := c.String("type")
		switch _type {
		case "json":
			file, err := os.Open(c.String("file"))
			if err != nil {
				return err
			}
			jb := jsonbook.New(file)
			jb.Authors = append(jb.Authors, "CodeTech BookMaker")
			book := &book.Book{
				Name:       jb.BookName,
				Authors:    jb.Authors,
				BookId:     uuid.New().String(),
				Cover:      coverFile,
				Chapters:   jb.GenChapters(),
				TempPath:   c.String("temp"),
				OutputPath: c.String("output"),
				Gen:        book.DefaultGen,
				Renderer:   book.NewRenderer(),
				PreChecker: book.DefaultPreChecker.Check,
			}
			return book.MakeMobi()
		case "txt":
			_, err := io.WriteString(c.App.Writer, "Not supported")
			return err
		default:
			return errors.New("unknown input file type", _type)
		}

	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
		},
		&cli.StringFlag{
			Name:    "cover",
			Aliases: []string{"c"},
			Value:   "./static/cover.jpg",
			Usage:   "the path of cover image",
		},
		&cli.StringFlag{
			Name:        "temp",
			Aliases:     []string{"t"},
			Value:       filepath.Join("./", uuid.New().String()),
			Usage:       "the path of temp files",
			DefaultText: "use a random uuid(v4)",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   filepath.Join("./output/"),
			Usage:   "the path of output files",
		},
		&cli.StringFlag{
			Name:  "type",
			Value: "json",
			Usage: "the input file type, json or txt",
		},
	},
}
