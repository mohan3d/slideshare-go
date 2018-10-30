package slideshare

import (
	"io"
)

// Downloader is the interface that wraps the basic Download method.
//
// Download downloads html of slideshare page and writes it to a given writer.
// It returns any error encountered that caused download or write to fail.
type Downloader interface {
	Download(w io.Writer, url string) error
}
