package book

import (
	"io"
	"path/filepath"
)

type StaticFile struct {
	Name    string
	Content io.Reader
}

func (s *StaticFile) MakeFile(path string) error {
	return globalRenderer.RenderFile(filepath.Join(path, s.Name), s.Content)
}
