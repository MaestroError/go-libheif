package libheif

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/strukturag/libheif/go/heif"
)

// saveAsJpeg encodes an image to JPEG format and saves it to a file.
// The quality should be an integer between 1 and 100, inclusive, higher is better.
//
// Example:
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	filename := "output.jpg"
// 	quality := 80
// 	if err := saveAsJpeg(img, filename, quality); err != nil {
// 		log.Fatal(err)
// 	}
func saveAsJpeg(img image.Image, filename string, quality int) error {
	// Validate input
	if img == nil {
		return fmt.Errorf("image is nil")
	}
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}
	if quality < 1 || quality > 100 {
		return fmt.Errorf("quality should be between 1 and 100")
	}

	// Encode the image to JPEG
	var out bytes.Buffer
	if err := jpeg.Encode(&out, img, &jpeg.Options{Quality: quality}); err != nil {
		return fmt.Errorf("could not encode image as JPEG: %w", err)
	}

	// Save the JPEG data to a file
	if err := ioutil.WriteFile(filename, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("could not save JPEG image as %s: %w", filename, err)
	}

	log.Printf("Image successfully written to %s\n", filename)
	return nil
}

// saveAsPng encodes an image to PNG format and saves it to a file.
//
// Example:
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	filename := "output.png"
// 	if err := saveAsPng(img, filename); err != nil {
// 		log.Fatal(err)
// 	}
func saveAsPng(img image.Image, filename string) error {
	// Validate input
	if img == nil {
		return fmt.Errorf("image is nil")
	}
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}

	// Encode the image to PNG
	var out bytes.Buffer
	if err := png.Encode(&out, img); err != nil {
		return fmt.Errorf("could not encode image as PNG: %w", err)
	}

	// Save the PNG data to a file
	if err := ioutil.WriteFile(filename, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("could not save PNG image as %s: %w", filename, err)
	}

	log.Printf("Image successfully written to %s\n", filename)
	return nil
}

// decodeHeifImage reads a file and decodes it into an image.
//
// This function opens the file specified by filename and decodes it into an image.
// It returns the image and the format of the image (like "jpeg", "png", "gif", "bmp" etc.).
// If an error occurs during opening the file or decoding the image, it is returned.
//
// Example:
// 	filename := "image.heic"
// 	img, format, err := decodeHeifImage(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// img is now an image.Image object and format is the format of the image.
func decodeHeifImage(filename string) (image.Image, string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		// The file couldn't be opened
		return nil, "", fmt.Errorf("could not open file %s: %v", filename, err)
	}
	// Ensure the file will be closed once we're done with it
	defer file.Close()

	// Decode the file into an image
	img, format, err := image.Decode(file)
	if err != nil {
		// The file couldn't be decoded
		return nil, "", fmt.Errorf("could not decode image: %v", err)
	}

	// Return the decoded image and its format
	return img, format, nil
}

// convertFileToHeif opens an image file, decodes it, and then converts
// it to the HEIF format.
//
// The new HEIF image is saved to a new file, the name of which is passed
// as a parameter. If an error occurs during the process, it is returned
// to the caller.
//
// Example:
// 	srcFile := "input.jpg"
// 	dstFile := "output.heif"
// 	if err := convertFileToHeif(srcFile, dstFile); err != nil {
// 		log.Fatal(err)
// 	}
func convertFileToHeif(filename string, newFileName string) error {
	// Validate input
	if filename == "" || newFileName == "" {
		return fmt.Errorf("input and output filenames must not be empty")
	}

	// Open the source file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	// Decode the image and get its format
	i, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Convert the image to HEIF and save it
	if err := SaveImageAsHeif(i, format, newFileName); err != nil {
		return fmt.Errorf("failed to convert image to HEIF and save it: %v", err)
	}

	return nil
}

// returnImageFromHeif decodes a HEIF file and returns the resulting image.
//
// It expects the filename of a HEIF file as input. If the file cannot be
// decoded or the decoded image is not in HEIF format, an error is returned.
//
// Example:
// 	filename := "image.heic"
// 	img, err := returnImageFromHeif(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
func returnImageFromHeif(filename string) (image.Image, error) {
	// Validate input
	if filename == "" {
		return nil, fmt.Errorf("filename must not be empty")
	}

	// Decode the image file
	img, format, err := decodeHeifImage(filename)
	if err != nil {
		return nil, fmt.Errorf("could not decode the image: %s; error: %v", filename, err)
	}

	// Check the format of the image (heic/heif/avif)
	if format != "heif" && format != "heic" && format != "avif" {
		return nil, fmt.Errorf("the image %s isn't in heif format, instead it is %s", filename, format)
	}

	// Return the decoded image
	return img, nil
}

// HeifToPng converts an image from HEIF format to PNG format.
//
// This function takes the path to the input HEIF image file and the path where
// the output PNG image file should be saved.
//
// Example:
// 	err := HeifToPng("input.heic", "output.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
func HeifToPng(heifImagePath string, newPngImagePath string) error {
	// Convert the HEIF image to the internal image representation.
	img, err := returnImageFromHeif(heifImagePath)
	if err != nil {
		// If there was an error, return it to the caller.
		return err
	}
	// Save the image in PNG format and return the result.
	return saveAsPng(img, newPngImagePath)
}

// HeifToJpeg converts an image from HEIF format to JPEG format.
//
// This function takes the path to the input HEIF image file, the path where
// the output JPEG image file should be saved, and the quality of the output JPEG image.
//
// Example:
// 	err := HeifToJpeg("input.heic", "output.jpeg", 80)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
func HeifToJpeg(heifImagePath string, newJpegImagePath string, quality int) error {
	// Convert the HEIF image to the internal image representation.
	img, err := returnImageFromHeif(heifImagePath)
	if err != nil {
		// If there was an error, return it to the caller.
		return err
	}
	// Save the image in JPEG format with the specified quality and return the result.
	return saveAsJpeg(img, newJpegImagePath, quality)
}

// ImageToHeif converts an image from JPEG or PNG format to HEIF format.
//
// This function takes the path to the input JPEG or PNG image file and the path where
// the output HEIF image file should be saved.
//
// Example:
// 	err := ImageToHeif("input.jpeg", "output.heic")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
func ImageToHeif(jpegOrPngImagePath string, newHeifImagePath string) error {
	// Convert the image to HEIF format and return the result.
	return convertFileToHeif(jpegOrPngImagePath, newHeifImagePath)
}

// SaveImageAsHeif encodes an image in HEIF format and saves it to a file.
//
// The image quality is set to 100, and lossless mode is enabled. The
// original image format is printed to the console.
//
// Example:
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	format := "png"
// 	newHeifImagePath := "output.heif"
// 	if err := SaveImageAsHeif(img, format, newHeifImagePath); err != nil {
// 		log.Fatal(err)
// 	}
func SaveImageAsHeif(i image.Image, format string, newHeifImagePath string) error {
	// Validate input
	if i == nil {
		return fmt.Errorf("image is nil")
	}
	if format == "" {
		return fmt.Errorf("format is empty")
	}
	if newHeifImagePath == "" {
		return fmt.Errorf("output path is empty")
	}

	// Encode the image in HEIF format
	const quality = 100
	ctx, err := heif.EncodeFromImage(i, heif.CompressionHEVC, quality, heif.LosslessModeEnabled, heif.LoggingLevelFull)
	if err != nil {
		return fmt.Errorf("failed to HEIF encode image: %w", err)
	}

	// Save the HEIF data to a file
	if err := ctx.WriteToFile(newHeifImagePath); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	log.Printf("Image was %s format, Written to %s\n", format, newHeifImagePath)
	return nil
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
