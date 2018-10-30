package slideshare

import (
	"io"
)

// Parser is the interface that wraps the basic Images method.
//
// Images extracts image urls from the given html.
// It returns a slice of strings and any error encountered that caused the parser to fail.
type Parser interface {
	Images(r io.Reader) ([]string, error)
}
