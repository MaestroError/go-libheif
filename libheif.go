package libheif

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/strukturag/libheif/go/heif"
)

func saveAsJpeg(img image.Image, filename string, quality int) {
	var out bytes.Buffer
	if err := jpeg.Encode(&out, img, &jpeg.Options{Quality: quality}); err != nil {
		fmt.Printf("Could not encode image as JPEG: %s\n", err)
	} else {
		if err := ioutil.WriteFile(filename, out.Bytes(), 0644); err != nil {
			fmt.Printf("Could not save JPEG image as %s: %s\n", filename, err)
		} else {
			fmt.Printf("Written to %s\n", filename)
		}
	}
}

func saveAsPng(img image.Image, filename string) {
	var out bytes.Buffer
	if err := png.Encode(&out, img); err != nil {
		fmt.Printf("Could not encode image as PNG: %s\n", err)
	} else {
		if err := ioutil.WriteFile(filename, out.Bytes(), 0644); err != nil {
			fmt.Printf("Could not save PNG image as %s: %s\n", filename, err)
		} else {
			fmt.Printf("Written to %s\n", filename)
		}
	}
}

func decodeHeifImage(filename string) (image.Image, string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not read file %s: %s\n", filename, err)
		return nil, "error"
	}
	defer file.Close()

	img, magic, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Could not decode image: %s\n", err)
		return nil, "error"
	}

	// fmt.Printf("Decoded image of type %s: %s\n", magic, img.Bounds())

	return img, magic
}

func convertImageToHeif(filename string, newFileName string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %v\n", file)
		return
	}
	defer file.Close()

	i, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("failed to decode image: %v\n", err)
		return
	}

	const quality = 100
	ctx, err := heif.EncodeFromImage(i, heif.CompressionHEVC, quality, heif.LosslessModeEnabled, heif.LoggingLevelFull)
	if err != nil {
		fmt.Printf("failed to heif encode image: %v\n", err)
		return
	}

	out := newFileName
	if err := ctx.WriteToFile(out); err != nil {
		fmt.Printf("failed to write to file: %v\n", err)
		return
	}
	fmt.Printf("Imaage was %s format, Written to %s\n", format, out)
}

func exampleHeifLowlevel(filename string) {
	fmt.Printf("Performing lowlevel conversion of %s\n", filename)
	c, err := heif.NewContext()
	if err != nil {
		fmt.Printf("Could not create context: %s\n", err)
		return
	}

	if err := c.ReadFromFile(filename); err != nil {
		fmt.Printf("Could not read file %s: %s\n", filename, err)
		return
	}

	nImages := c.GetNumberOfTopLevelImages()
	fmt.Printf("Number of top level images: %v\n", nImages)

	ids := c.GetListOfTopLevelImageIDs()
	fmt.Printf("List of top level image IDs: %#v\n", ids)

	if pID, err := c.GetPrimaryImageID(); err != nil {
		fmt.Printf("Could not get primary image id: %s\n", err)
	} else {
		fmt.Printf("Primary image: %v\n", pID)
	}

	handle, err := c.GetPrimaryImageHandle()
	if err != nil {
		fmt.Printf("Could not get primary image: %s\n", err)
		return
	}

	fmt.Printf("Image size: %v × %v\n", handle.GetWidth(), handle.GetHeight())

	img, err := handle.DecodeImage(heif.ColorspaceUndefined, heif.ChromaUndefined, nil)
	if err != nil {
		fmt.Printf("Could not decode image: %s\n", err)
	} else if i, err := img.GetImage(); err != nil {
		fmt.Printf("Could not get image: %s\n", err)
	} else {
		fmt.Printf("Rectangle: %v\n", i.Bounds())

		ext := filepath.Ext(filename)
		outFilename := filename[0:len(filename)-len(ext)] + "_lowlevel.png"
		saveAsPng(i, outFilename)
	}
}

func returnImageFromHeif(filename string) image.Image {
	img, format := decodeHeifImage(filename)
	if format == "error" {
		fmt.Println("Could not decode the image: %s ; error returned\n", filename)
	}
	if format != "heif" && format != "heic" {
		fmt.Println("The image %s isn't heif format, instead it is %s\n", filename, format)
	}
	return img
}

func HeifToPng(heifImagePath string, newPngImagePath string) {
	img := returnImageFromHeif(heifImagePath)
	if img == nil {
		return
	}
	saveAsPng(img, newPngImagePath)
}

func HeifToJpeg(heifImagePath string, newJpegImagePath string) {
	img := returnImageFromHeif(heifImagePath)
	if img == nil {
		return
	}
	saveAsJpeg(img, newJpegImagePath, 100)
}
