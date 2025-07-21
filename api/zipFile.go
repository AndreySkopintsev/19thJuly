package api

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// useful link https://earthly.dev/blog/golang-zip-files/

func CreateAnArchive(directory string) error {
	fileNames, err := GetFilesInADirectory(directory)
	if err != nil {
		return err
	}
	zipFile, err := os.Create(directory + "/" + "archive.zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipw := zip.NewWriter(zipFile)

	for _, fileName := range fileNames {
		fileToZip, err := os.Open(directory + "/" + fileName)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		w1, err := zipw.Create(fileName)
		if err != nil {
			fmt.Printf("couldnt create in zip with error %s", err.Error())
			return err
		}

		if _, err := io.Copy(w1, fileToZip); err != nil {
			fmt.Printf("couldnt crcopy file with error %s", err.Error())
			return err
		}

	}

	zipw.Close()
	return nil
}

func GetFilesInADirectory(filePath string) ([]string, error) {
	entries, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	fileNames := []string{}
	for _, entry := range entries {
		fileNames = append(fileNames, entry.Name())
	}
	return fileNames, nil
}
