package book

import (
	"github.com/c-jimin/BookMaker/errors"
	"io"
	"os"
	"os/exec"
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
	if _, err := os.Stat(genPath); err != nil {
		return "", err
	}
	return genPath, nil
}

func DefaultGen(book *Book) error {
	genPath, err := getGenPath()
	if err != nil {
		return err
	}
	cmd := exec.Command(genPath, "-dont_append_source", book.TempPath+book.Name+".opf")
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
