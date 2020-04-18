package main

import (
	"github.com/c-jimin/http-package/request"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func combine(str ...string) string {
	return strings.Join(str, "")
}

func downloadFile(path, url string) error {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return nil
	}
	resp := request.Request{
		Method: "GET",
		URL:    url,
	}.SendWithClient(client)
	if resp.HasError() {
		return resp.Error
	}
	defer resp.Close()
	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, 0775); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp)
	return nil
}
