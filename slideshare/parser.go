package slideshare

import (
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Parser is the interface that wraps the basic Images method.
//
// Images extracts image urls from the given html.
// It returns a slice of strings and any error encountered that caused the parser to fail.
type Parser interface {
	Images(r io.Reader, q Quality) ([]string, error)
}

// Quality represents image quality
type Quality int

const (
	// QualityFull best quality.
	QualityFull Quality = iota + 1

	// QualityNormal normal quality.
	QualityNormal

	// QualitySmall worst quality.
	QualitySmall
)

const (
	slideSelector          = "img.slide_image"
	imageURLSelectorFull   = "data-full"
	imageURLSelectorNormal = "data-normal"
	imageURLSelectorSmall  = "data-small"
)

var errNoImages = errors.New("parsing returns no images or error")

func imageSelector(q Quality) string {
	var selector string

	switch q {
	case QualityFull:
		selector = imageURLSelectorFull
	case QualityNormal:
		selector = imageURLSelectorNormal
	case QualitySmall:
		selector = imageURLSelectorSmall
	}
	return selector
}

type defaultParser struct {
}

func (p *defaultParser) Images(r io.Reader, q Quality) ([]string, error) {
	d, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		return nil, err
	}
	var urls []string

	d.Find(slideSelector).Each(func(i int, s *goquery.Selection) {
		selector := imageSelector(q)
		if url, exists := s.Attr(selector); exists {
			urls = append(urls, url)
		}
	})
	if len(urls) == 0 {
		return nil, errNoImages
	}
	return urls, nil
}

// DefaultParser used when no implementation
// is provided for Parser interface.
var DefaultParser = &defaultParser{}
