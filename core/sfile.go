package core

import "io"

type StaticFile struct {
	Name    string
	Content io.Reader
}

func (s *StaticFile) MakeFile(path string) error {
	return globalRenderer.RenderFile(path, s.Name, s.Content)
}
