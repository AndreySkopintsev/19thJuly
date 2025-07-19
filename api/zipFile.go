package api

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

func CreateAnArchive(directory string) error {
	fileNames, err := GetFilesInADirectory(directory)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, fileName := range fileNames {
		f, err := w.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(directory + "/" + fileName))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.WriteFile(directory+"/"+"data.zip", buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

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
