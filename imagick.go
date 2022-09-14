package main

import (
	exifremove "github.com/scottleedavis/go-exif-remove"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gographics/imagick.v3/imagick"
)

const filePathMagic = `/Users/aleksandr/GolandProjects/test_libimagequant/test_images/IMG_2775.jpg`

func main() {
	var err error

	imagick.Initialize()
	defer imagick.Terminate()

	rawImg, err := os.ReadFile(filePathMagic)
	if err != nil {
		log.Fatalf("can't read file %s: %s", filepath.Base(filePathMagic), err)
	}

	rawImg, err = exifremove.Remove(rawImg)
	if err != nil {
		log.Fatalf("fail remove exif data: %s", err)
	}

	mw := imagick.NewMagickWand()

	err = mw.ReadImageBlob(rawImg)
	if err != nil {
		log.Fatalf("can't read image from blob: %s", err)
	}

	err = mw.StripImage()
	if err != nil {
		log.Fatalf("can't strip image from all profiles and comments: %s", err)
	}

	err = mw.SetImageCompressionQuality(55)
	if err != nil {
		log.Fatalf("can't set compression: %s", err)
	}

	err = mw.BlurImage(0, 0.75)
	if err != nil {
		log.Fatalf("can't blur image: %s", err)
	}

	err = mw.SetImageInterlaceScheme(imagick.INTERLACE_PLANE)
	if err != nil {
		log.Fatalf("can't set plane interlace scheme: %s", err)
	}

	err = mw.SetImageFormat("JPEG")
	if err != nil {
		log.Fatalf("can't set image format: %s", err)
	}

	err = mw.SetColorspace(imagick.COLORSPACE_SCRGB)
	if err != nil {
		log.Fatalf("can't set color space: %s", err)
	}

	compressedImg := mw.GetImageBlob()

	err = os.WriteFile(
		filepath.Dir(filePathMagic)+"/magic_compressed_"+filepath.Base(filePathMagic),
		compressedImg,
		0644)
	if err != nil {
		log.Fatalf("can't write compressed img to disk: %s", err)
	}

	log.Printf("Done. Raw size: %d, compressed size: %d, win: %.1f%%",
		len(rawImg), len(compressedImg), (float32(len(rawImg))-float32(len(compressedImg)))/float32(len(rawImg))*100)
}
