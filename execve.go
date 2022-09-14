package main

import (
	"fmt"
	"log"
	"path/filepath"
	"syscall"
	"time"
)

const (
	cjpegPath  = `/Users/aleksandr/GolandProjects/test_libimagequant/cjpeg`
	rawImgPath = `/Users/aleksandr/GolandProjects/test_libimagequant/test_images/IMG_2775.JPG`
	quality    = 65
)

func main() {
	resultImgPath := fmt.Sprintf("%s/%d_%s", filepath.Dir(rawImgPath), time.Now().UnixNano(), filepath.Base(rawImgPath))

	argvArr := []string{"-quality " + fmt.Sprintf("%d", quality), "-outfile " + resultImgPath, rawImgPath}
	log.Printf("%v", argvArr)

	err := syscall.Exec(cjpegPath, argvArr, nil)
	if err != nil {
		log.Fatalf("can't compress %s to %s with %d quality: %s",
			rawImgPath, resultImgPath, quality, err)
	}

	log.Printf("%s created with quality %d", filepath.Base(resultImgPath), quality)
}
