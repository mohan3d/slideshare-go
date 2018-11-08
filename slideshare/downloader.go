package slideshare

import (
	"io"
	"net/http"
)

// Downloader is the interface that wraps the basic Download method.
//
// Download downloads html of slideshare page and writes it to a given writer.
// It returns any error encountered that caused download or write to fail.
type Downloader interface {
	Fetch(w io.Writer, url string) error
}

type defaultDownloader struct{}

func (d *defaultDownloader) Fetch(w io.Writer, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	return err
}

// DefaultDownloader used when no implementation
// is provided for Downloader interface.
var DefaultDownloader = &defaultDownloader{}
