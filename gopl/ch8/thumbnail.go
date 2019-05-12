package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//将一个jpeg图像格式缩放

//缩放图像
func Image(src image.Image) image.Image {
	//计算缩略图大小，保持纵横比
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y

	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	decode, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(decode)
	return jpeg.Encode(w, dst, nil)
}

func ImageFile2(outfile, infile string) (err error) {
	file, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer file.Close()
	create, err := os.Create(infile)
	if err != nil {
		return err
	}
	if err := ImageStream(file, create); err != nil {
		create.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return create.Close()
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		thumb, err := ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
