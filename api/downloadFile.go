package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var ErrExtension = errors.New("invalid extension")

// https://gophercoding.com/download-a-file/ implementation
func DownloadFile(filename string, url, directoryName string) error {
	directoryName = strings.Trim(directoryName, "./")
	currentPath, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}
	err = os.MkdirAll(currentPath+"/"+directoryName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}
	fmt.Printf("current path %s\n", currentPath)
	// Get the data
	fmt.Println(url + "\n")
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
	out, err := os.Create(currentPath + "/" + directoryName + "/" + filename + extension)
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
