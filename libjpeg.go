package main

import (
	"bufio"
	"bytes"
	"github.com/pixiv/go-libjpeg/jpeg"
	exifremove "github.com/scottleedavis/go-exif-remove"
	"log"
	"os"
	"path/filepath"
)

const filePathLibJPEG = `/app/test_images/IMG_2775.jpg`

func main() {
	rawImg, err := os.ReadFile(filePathLibJPEG)
	if err != nil {
		log.Fatalf("can't read file %s: %s", filepath.Base(filePathLibJPEG), err)
	}

	rawImg, err = exifremove.Remove(rawImg)
	if err != nil {
		log.Fatalf("fail remove exif data: %s", err)
	}

	img, err := jpeg.Decode(bytes.NewReader(rawImg), &jpeg.DecoderOptions{})
	if err != nil {
		log.Fatalf("Decode returns error: %v\n", err)
	}

	var buf bytes.Buffer

	err = jpeg.Encode(bufio.NewWriter(&buf), img, &jpeg.EncoderOptions{
		Quality:         50,
		OptimizeCoding:  true,
		ProgressiveMode: true,
		DCTMethod:       jpeg.DCTISlow,
	})
	if err != nil {
		log.Fatalf("Encode returns error: %v\n", err)
	}
	compressedImg := buf.Bytes()

	err = os.WriteFile(
		filepath.Dir(filePathLibJPEG)+"/libjpeg_compressed_"+filepath.Base(filePathLibJPEG),
		compressedImg,
		0644)
	if err != nil {
		log.Fatalf("can't write compressed img to disk: %s", err)
	}

	log.Printf("Done. Raw size: %d, compressed size: %d, win: %.1f%%",
		len(rawImg), len(compressedImg), (float32(len(rawImg))-float32(len(compressedImg)))/float32(len(rawImg))*100)
}
