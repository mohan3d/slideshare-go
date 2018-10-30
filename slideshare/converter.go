package slideshare

import "io"

// Converter is the interface that wraps image conversion methods.
//
// Reset resets the internal state of the parser.
// It may be empty based on the implementation.
//
// AddImage append new image to the pdf.
// It returns any error encountered that caused save tp fail.
//
// Save writes pdf content to an io.Writer.
// It returns any error encountered that caused save to fail.
type Converter interface {
	Reset()
	AddImage(r io.Reader, name string) error
	Save(w io.Writer) error
}
