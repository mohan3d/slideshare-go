package slideshare

import (
	"io"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

// Converter is the interface that wraps image conversion methods.
//
// Reset resets the internal state of the parser.
// It may be empty based on the implementation.
//
// AddImage append new image to the pdf.
// It returns any error encountered that caused save to fail.
//
// Save writes pdf content to an io.Writer.
// It returns any error encountered that caused save to fail.
type Converter interface {
	Reset()
	AddImage(r io.Reader, name string) error
	Save(w io.Writer) error
}

func imageExtension(name string) string {
	ext := filepath.Ext(name)
	return ext[1:]
}

// defaultConverter defaultConverter
type defaultConverter struct {
	pdf *gofpdf.Fpdf
	opt gofpdf.ImageOptions
}

func (c *defaultConverter) Reset() {
	c.pdf = gofpdf.New("L", "", "", "")
}

func (c *defaultConverter) AddImage(r io.Reader, name string) error {
	c.pdf.AddPage()
	c.pdf.RegisterImageReader(name, imageExtension(name), r)

	w, h := c.pdf.GetPageSize()
	c.pdf.ImageOptions(name, 0, 0, w, h, false, c.opt, 0, "")

	if c.pdf.Err() {
		return c.pdf.Error()
	}
	return nil
}

func (c *defaultConverter) Save(w io.Writer) error {
	return c.pdf.Output(w)
}

func newDefaultConverter() *defaultConverter {
	c := new(defaultConverter)
	c.Reset()
	return c
}

// DefaultConverter used when no implementation
// is provided for Converter interface.
var DefaultConverter = newDefaultConverter()
