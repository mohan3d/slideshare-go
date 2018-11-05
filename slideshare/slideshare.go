package slideshare

import (
	"bytes"
	"fmt"
	"io"
)

// PDFDownloader manages the process of pdf creation.
// It uses downloader to download slideshare pages html and slide images
// parser to extract images urls from slideshare pages
// and converter to convert images into a pdf.
type PDFDownloader struct {
	downloader Downloader
	parser     Parser
	converter  Converter
}

// Download downloads pdf from slideshare url with given quality
// and writes it to w.
// It returns any error encountered in download, parse or convert.
func (pd *PDFDownloader) Download(url string, quality Quality, w io.Writer) error {
	var html bytes.Buffer
	err := pd.downloader.Fetch(&html, url)

	if err != nil {
		return err
	}
	imageURLs, err := pd.parser.Images(&html, quality)

	if err != nil {
		return err
	}
	// reset before creating a new pdf.
	pd.converter.Reset()

	for i, imageURL := range imageURLs {
		var img bytes.Buffer
		err = pd.downloader.Fetch(&img, imageURL)
		if err != nil {
			return err
		}
		err = pd.converter.AddImage(&img, fmt.Sprintf("%d.jpg", i))
		if err != nil {
			return err
		}
	}
	return pd.converter.Save(w)
}

// NewSlideshareDownloader returns a new instance of PDFDownloader
// created from the given downloader, parser and converter.
func NewSlideshareDownloader(downloader Downloader, parser Parser, converter Converter) *PDFDownloader {
	return &PDFDownloader{
		downloader: downloader,
		parser:     parser,
		converter:  converter,
	}
}

// DefaultSlideshareDownloader used if the default implementation
// of downloader, parser and converter is sufficient.
var DefaultSlideshareDownloader = NewSlideshareDownloader(DefaultDownloader, DefaultParser, DefaultConverter)
