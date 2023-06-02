# go-libheif
 GoLang wrapper for the libheif library, providing easy-to-use APIs for HEIC to JPEG/PNG conversions and vice versa. (Also provides support for AVIF to JPEG/PNG conversions)      
 This package was developed to support 

### Pre-requisites
You need to install [libheif](https://github.com/strukturag/libheif) before using this module. You can check the [strukturag/libheif](https://github.com/strukturag/libheif) for installation instructions, but as I have found, the easiest way for me was to use [brew](https://brew.sh/):
```bash
brew install cmake make pkg-config x265 libde265 libjpeg libtool
brew install libheif
```
*Note: Pay attention to the brew's recommendation while installing*    

### Installation
```bash
go get github.com/MaestroError/go-libheif
```

### Usage
This library provides functions to convert images between HEIC format and other common formats such as JPEG and PNG.       
          
To convert an image from HEIC / HEIF / AVIF to JPEG or PNG:
```go
package main

import (
	"log"

	"github.com/MaestroError/go-libheif"
)

func main() {
	err := go_libheif.HeifToJpeg("input.heic", "output.jpeg", 80)
	if err != nil {
		log.Fatal(err)
	}
	
	err = go_libheif.HeifToPng("input.heic", "output.png")
	if err != nil {
		log.Fatal(err)
	}
}

```

To convert an image from JPEG or PNG to HEIC:
```go
package main

import (
	"log"

	"github.com/MaestroError/go-libheif"
)

func main() {
	err := go_libheif.ImageToHeif("input.jpeg", "output.heic")
	if err != nil {
		log.Fatal(err)
	}
}

```
*Note: It finds hard to convert some jpeg and png files to heic, see libheif_test.go:14 for details*       
           
To save an image as HEIC:
```go
package main

import (
	"image"
	"log"

	"github.com/MaestroError/go-libheif"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	err := go_libheif.SaveImageAsHeif(img, "png", "output.heic")
	if err != nil {
		log.Fatal(err)
	}
}

```
*Note: Quality for **HeifToJpeg** function and image format for **SaveImageAsHeif** function should be provided. The quality value ranges from 1 to 100 inclusive, higher values meaning better quality. The format for **SaveImageAsHeif** is the format of the original image.*       

Please consult the GoDoc documentation for more detailed information about the provided functions and their usage.          


### Credits
Thanks to @strukturag and @farindk (Dirk Farin) for his work on the [libheif](https://github.com/strukturag/libheif) library 🙏

##### ToDo
- Write usage documentation +
- Write guide article for module
- Add contribution section in readme
- Implement the module in [php-heic-to-jpg](https://github.com/MaestroError/php-heic-to-jpg) +
- Add credits section in readme and update in php-heic-to-jpg +