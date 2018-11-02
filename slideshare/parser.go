package slideshare

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Parser is the interface that wraps the basic Images method.
//
// Images extracts image urls from the given html.
// It returns a slice of strings and any error encountered that caused the parser to fail.
type Parser interface {
	Images(r io.Reader) ([]string, error)
}

const (
	slideSelector    = "img.slide_image"
	imageURLSelector = "data-full"
)

type defaultParser struct {
}

func (p *defaultParser) Images(r io.Reader) ([]string, error) {
	d, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		return nil, err
	}
	var urls []string

	d.Find(slideSelector).Each(func(i int, s *goquery.Selection) {
		if url, exists := s.Attr(imageURLSelector); exists {
			urls = append(urls, url)
		}
	})
	return urls, nil
}

// DefaultParser used when no implementation
// is provided for Parser interface.
var DefaultParser = &defaultParser{}
