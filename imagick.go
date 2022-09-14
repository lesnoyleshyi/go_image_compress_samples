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

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	//err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	//if err != nil {
	//	log.Fatalf("can't resize image: %s", err)
	//}

	err = mw.SetImageCompressionQuality(75)
	if err != nil {
		log.Fatalf("can't set compression: %s", err)
	}

	err = mw.SetImageFormat("JPEG")
	if err != nil {
		log.Fatalf("can't set image format: %s", err)
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
