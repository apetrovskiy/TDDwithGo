package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
