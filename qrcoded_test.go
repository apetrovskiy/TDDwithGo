package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGenratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")

	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
