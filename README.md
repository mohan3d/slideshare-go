# slideshare-go
[![GoDoc](https://godoc.org/github.com/mohan3d/slideshare-go?status.svg)](https://godoc.org/github.com/mohan3d/slideshare-go/slideshare)
[![Go Report Card](https://goreportcard.com/badge/github.com/mohan3d/slideshare-go)](https://goreportcard.com/report/github.com/mohan3d/slideshare-go)    
Golang package to download [slideshare](https://www.slideshare.net/) slides to pdf files.

## Installation
```bash
$ go get github.com/mohan3d/slideshare-go
```

## Usage
```sh
$ slideshare-go -url <slideshare_url> -path <path>
$ slideshare-go -url <slideshare_url> -path <path> -quality <quality>
$ slideshare-go --help
```

## Examples
```sh
# download highest quality of 
# https://www.slideshare.net/DeloitteUS/the-hospital-of-the-future-81817523
# and save it to hospital.pdf 
$ slideshare-go -url=https://www.slideshare.net/DeloitteUS/the-hospital-of-the-future-81817523 -path=./hospital.pdf -quality=high

# download normal quality of 
# https://www.slideshare.net/jeanbaptiste.dumont/the-ai-rush
# and save it to ai_rush.pdf 
$ slideshare-go -url=https://www.slideshare.net/jeanbaptiste.dumont/the-ai-rush -path=./ai_rush.pdf -quality=normal

# download lowest quality of 
# https://www.slideshare.net/carologic/ai-and-machine-learning-demystified-by-carol-smith-at-midwest-ux-2017
# and save it to ai.pdf 
$ slideshare-go -url=https://www.slideshare.net/carologic/ai-and-machine-learning-demystified-by-carol-smith-at-midwest-ux-2017 -path=./ai.pdf -quality=low
```
