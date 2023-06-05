# go-libheif
 GoLang wrapper for the libheif library, providing easy-to-use APIs for HEIC to JPEG/PNG conversions and vice versa (Also provides support for AVIF to JPEG/PNG conversions).      
 This package was originally developed to support the [php-heic-to-jpg](https://github.com/MaestroError/php-heic-to-jpg) package, which had [problems](https://github.com/MaestroError/php-heic-to-jpg/issues/15) while converting some HEIF images.   

#### Implementation:
- For a swift start, consider browsing our [quickstart guide](https://medium.com/@revaz.gh/using-the-go-libheif-module-for-converting-images-between-the-heic-format-and-other-popular-formats-e7829165c368) on Medium.
- If you're primarily interested in a hassle-free image conversion tool, we have already integrated this module into a Command Line Interface (CLI) application and a docker image, available [here](https://github.com/MaestroError/heif-converter-image).

### Prerequisites
You need to install [libheif](https://github.com/strukturag/libheif) before using this module. You can check the [strukturag/libheif](https://github.com/strukturag/libheif) for installation instructions, but as I have found, the easiest way for me was to use [brew](https://brew.sh/):
```bash
brew install cmake make pkg-config x265 libde265 libjpeg libtool
brew install libheif
```
*Note: Pay attention to the brew's recommendation while installing*    

You can find installation scripts in [heif-converter-image](https://github.com/MaestroError/heif-converter-image) repository with names:
- install-libheif.sh
- install-libheif-macos.sh
- install-libheif-windows.bat

*If the installation process seems a bit challenging, we recommend to consider using the [heif-converter-image](https://github.com/MaestroError/heif-converter-image) which offers a ready-to-use docker image.*

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
	err := libheif.HeifToJpeg("input.heic", "output.jpeg", 80)
	if err != nil {
		log.Fatal(err)
	}
	
	err = libheif.HeifToPng("input.heic", "output.png")
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
	err := libheif.ImageToHeif("input.jpeg", "output.heic")
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
	err := libheif.SaveImageAsHeif(img, "png", "output.heic")
	if err != nil {
		log.Fatal(err)
	}
}

```
*Note: Quality for **HeifToJpeg** function and image format for **SaveImageAsHeif** function should be provided. The quality value ranges from 1 to 100 inclusive, higher values meaning better quality. The format for **SaveImageAsHeif** is the format of the original image.*       

Please consult the GoDoc [documentation](https://pkg.go.dev/github.com/MaestroError/go-libheif) for more detailed information about the provided functions and their usage.          

### Contributions
Contributions to `go-libheif` are wholeheartedly encouraged! We believe in the power of the open source community to build the most robust and accessible projects. Whether you are a seasoned Go developer or just getting started, your perspective and efforts can help improve this project for everyone.

Here are a few ways you might contribute:

- Bug Fixes: If you encounter a problem with go-libheif, please open an issue in the GitHub repository. Even better, if you can solve the issue, we welcome pull requests.
- New Features: Interested in expanding the capabilities of go-libheif? Please open an issue to discuss your idea before starting on the code. Once we've agreed on an approach, you can submit a pull request with your new feature.
- Documentation: Clear documentation is vital to any project. If you can clarify our instructions or add helpful examples, we would appreciate your input.

Here are the steps for contributing code:

- Fork the go-libheif repository and clone it to your local machine.
- Create a new branch for your changes.
- Make your changes and push them to your fork.
- Open a pull request from your fork's branch to the go-libheif main branch.

Please remember to be as detailed as possible in your pull request descriptions to help the maintainers understand your changes.
We look forward to seeing what you contribute! Together, we can make go-libheif even better.

### Credits
Thanks to @strukturag and @farindk (Dirk Farin) for his work on the [libheif](https://github.com/strukturag/libheif) library 🙏