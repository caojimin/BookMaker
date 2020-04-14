package book

import (
	"bytes"
)

var globalRenderer = NewRenderer()

func DefaultPretreatment(c *Chapter) error {
	buffer := bytes.NewBuffer(nil)
	if err := globalRenderer.Reducer.Minify("text/xml", buffer, c.Content); err != nil {
		return err
	}
	c.Content = buffer
	return nil
}
