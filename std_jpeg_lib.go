package main

import (
	"bufio"
	"bytes"
	exifremove "github.com/scottleedavis/go-exif-remove"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

const filePath = `/Users/aleksandr/GolandProjects/test_libimagequant/test_images/IMG_2775.jpg`

//const speed = 10

func main() {
	rawImg, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("can't read file %s: %s", filepath.Base(filePath), err)
	}

	rawImg, err = exifremove.Remove(rawImg)
	if err != nil {
		log.Fatalf("fail remove exif data: %s", err)
	}
	//compressedImg, err := imagequant.Crush(rawImg, speed, png.BestCompression)
	//if err != nil {
	//	log.Fatalf("can't compress file: %s", err)
	//}
	imageData, err := jpeg.Decode(bytes.NewReader(rawImg))
	if err != nil {
		log.Fatalf("can't decode image to image.Image: %s", err)
	}

	//newImage := resize.Resize(width, 0, imageData, resize.Lanczos3)
	newImage := imageData

	var buf bytes.Buffer

	err = jpeg.Encode(bufio.NewWriter(&buf), newImage, &jpeg.Options{Quality: 50})
	if err != nil {
		log.Fatalf("can't decode jpg: %s", err)
	}
	compressedImg := buf.Bytes()

	err = os.WriteFile(
		filepath.Dir(filePath)+"/compressed_"+filepath.Base(filePath),
		compressedImg,
		0644)
	if err != nil {
		log.Fatalf("can't write compressed img to disk: %s", err)
	}

	log.Printf("Done. Raw size: %d, compressed size: %d, win: %.1f%%",
		len(rawImg), len(compressedImg), (float32(len(rawImg))-float32(len(compressedImg)))/float32(len(rawImg))*100)
}
