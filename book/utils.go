package book

import (
	"bytes"
	"io/ioutil"
)

var globalRenderer = NewRenderer()

func DefaultPretreatment(c *Chapter) error {
	buffer := bytes.NewBuffer(nil)
	if err := globalRenderer.Reducer.Minify("text/xml", buffer, c.Content); err != nil {
		return err
	}
	c.Content = ioutil.NopCloser(buffer)
	return nil
}
