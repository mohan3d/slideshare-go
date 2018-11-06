package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/mohan3d/slideshare-go/slideshare"
)

func selectQuality(q string) slideshare.Quality {
	// defaults to high quality
	quality := slideshare.QualityFull

	switch q {
	case "normal":
		quality = slideshare.QualityNormal
	case "low":
		quality = slideshare.QualitySmall
	}
	return quality
}

func download(url, quality, path string) {
	var buf bytes.Buffer

	err := slideshare.DefaultSlideshareDownloader.Download(
		url,
		selectQuality(quality),
		&buf,
	)
	if err != nil {
		panic(err)
	}
	output, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	_, err = io.Copy(output, &buf)
	if err != nil {
		panic(err)
	}
}

func main() {
	url := flag.String("url", "", "slideshare url to be downloaded")
	path := flag.String("path", "", "path to save the pdf")
	quality := flag.String("quality", "high", "quality of pdf it should be \"high\", \"normal\" or \"low\"")

	flag.Parse()

	if !(*url != "" && *path != "") {
		fmt.Println("A valid slideshare url and save path must be provided\n")
		flag.Usage()
		os.Exit(1)
	}
	download(*url, *quality, *path)
}
