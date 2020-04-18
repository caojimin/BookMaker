package main

import (
	"github.com/c-jimin/BookMaker/errors"
	"github.com/c-jimin/http-package/client"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"runtime"
)

var client = httpClient.New()

func InitCover() error {
	return downloadFile("./static/cover.jpg", "https://github.com/c-jimin/BookMaker/raw/master/static/cover.jpg")
}

func InitGen() error {
	var ospath, filename string
	switch runtime.GOOS {
	case "windows":
		ospath, filename = "windows", "kindlegen.exe"
	case "darwin":
		ospath, filename = "mac", "kindlegen"
	case "linux":
		ospath, filename = "linux", "kindlegen"
	default:
		return errors.New("unknown os", runtime.GOOS)
	}
	return downloadFile(
		filepath.Join("./bin", ospath, filename),
		combine("https://github.com/c-jimin/BookMaker/raw/master/bin/", ospath, "/", filename),
	)
}

var cmdInit = &cli.Command{
	Name:  "init",
	Usage: "init the static file",
	Action: func(c *cli.Context) error {
		proxy := c.String("proxy")
		if proxy != "" {
			if err := client.SetProxy(proxy); err != nil {
				return err
			}
		}
		if err := InitCover(); err != nil {
			return err
		}
		if err := InitGen(); err != nil {
			return err
		}
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "proxy",
			Aliases: []string{"p"},
		},
	},
}
