package book

import (
	"github.com/c-jimin/BookMaker/errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getGenPath() (string, error) {
	var genPath string
	switch runtime.GOOS {
	case "windows":
		genPath = "./bin/windows/kindlegen.exe"
	case "darwin":
		genPath = "./bin/mac/kindlegen"
	case "linux":
		genPath = "./bin/linux/kindlegen"
	default:
		return "", errors.New("unknown os", runtime.GOOS)
	}
	fp, err := filepath.Abs(genPath)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(fp); err != nil {
		return "", err
	}
	return fp, nil
}

func DefaultGen(book *Book) error {
	genPath, err := getGenPath()
	if err != nil {
		return err
	}
	fp, err := filepath.Abs(filepath.Join(book.TempPath, book.Name+".opf"))
	if err != nil {
		return err
	}
	cmd := exec.Command(genPath, "-dont_append_source", fp)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	defer stdout.Close()
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	return cmd.Wait()
}
