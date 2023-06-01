# go-libheif
 GoLang wrapper for the libheif library, providing easy-to-use APIs for HEIC to JPEG/PNG conversions and vice versa.

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
...

##### ToDo
- Refactor the functions to return errors instead +
- Add documentation and examples for each function +
- Write tests at least for the Public functions +
- Write usage documentation
- Write guide article for module