package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var ErrExtension = errors.New("invalid extension")

// https://gophercoding.com/download-a-file/ implementation
func DownloadFile(filename string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	extension, err := GetFileExtension(url)
	if err != nil {
		return err
	}
	out, err := os.Create("./downloads/" + filename + extension)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// GetFileExtansion gets file extension from URl
func GetFileExtension(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", ErrExtension
	}

	filePath := parsedUrl.Path
	return filepath.Ext(filePath), nil
}
