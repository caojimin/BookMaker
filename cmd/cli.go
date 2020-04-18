package main

import (
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "bookmaker",
		Usage: "A util to make a book format mobi or ePub",
		Commands: []*cli.Command{
			cmdInit,
			cmdMake,
			{
				Name:  "server",
				Usage: "http server",
				Action: func(c *cli.Context) error {
					_, err := io.WriteString(c.App.Writer, "Not supported")
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
