package main

import (
	"bytes"
	"errors"
	"github.com/dailymotion/allure-go"
	"github.com/joho/godotenv"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	dotenv := goDotEnvVariable("ALLURE_RESULTS_PATH")
	allure.Test(t, allure.Action(func() {
		buffer := new(bytes.Buffer)
		GenerateQRCode(buffer, "555-2368", Version(1))

		if buffer.Len() == 0 {
			t.Errorf("No QRCode generated")
		}
		_, err := png.Decode(buffer)

		if err != nil {
			t.Errorf("Generated QRCode is not a PNG: %s", err)
		}
	}))
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2368", Version(1))

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range table {
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d",
				test.version, test.expected, size)
		}
	}
}

func TestGenerateQRCode(t *testing.T) {
	type args struct {
		code    string
		version Version
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := GenerateQRCode(w, tt.args.code, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("GenerateQRCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GenerateQRCode() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("variables.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
