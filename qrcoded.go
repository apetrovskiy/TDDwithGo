package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", Version(1))
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := 4*int(version) + 17
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

type Version int8
