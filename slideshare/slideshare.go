package slideshare

// PDFDownloader manages the process of pdf creation.
// It uses downloader to download slideshare pages html and slide images
// parser to extract images urls from slideshare pages
// and converter to convert images into a pdf.
type PDFDownloader struct {
	downloader Downloader
	parser     Parser
	converter  Converter
}
