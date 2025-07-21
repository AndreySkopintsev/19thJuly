package api

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	TestFilesPath = "./testFiles"
)

func TestDownloadJPEG(t *testing.T) {
	t.Run("donwloading a png", func(t *testing.T) {
		url := "https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750"

		err := DownloadFile("saveas", url, TestFilesPath)
		if err != nil {
			fmt.Println("Error downloading file: ", err)
			return
		}

		fmt.Println("Downloaded: " + url)

	})
}

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		Name              string
		RawURL            string
		ExpectedExtension string
		ExpectedError     error
	}{
		{
			Name:              "link to a png file",
			RawURL:            "https://gophercoding.com/img/logo-original.png",
			ExpectedExtension: ".png",
			ExpectedError:     nil,
		},
		{
			Name:              "link to a jpg file",
			RawURL:            "https://upload.wikimedia.org/wikipedia/commons/f/ff/Wikipedia_logo_593.jpg?20060603094750",
			ExpectedExtension: ".jpg",
			ExpectedError:     nil,
		},
		{
			Name:              "link to a pdf file",
			RawURL:            "https://www.antennahouse.com/hubfs/xsl-fo-sample/pdf/basic-link-1.pdf",
			ExpectedExtension: ".pdf",
			ExpectedError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := GetFileExtension(tt.RawURL)
			if err != nil {
				if !errors.Is(err, tt.ExpectedError) {
					t.Fatalf("expected error was %+v, got %+v err instead", tt.ExpectedError, err)
				}
			}

			if got != tt.ExpectedExtension {
				t.Fatalf("expected extension %s, got %s", tt.ExpectedExtension, got)
			}
		})
	}
}

func TestGetFileNames(t *testing.T) {
	tests := []struct {
		Name              string
		ExpectedFileNames []string
		FilesToCreate     []string
	}{
		{
			Name:              "create 3 txt files, expect to get 3 file names",
			ExpectedFileNames: []string{"test1.txt", "test2.txt", "test3.txt"},
			FilesToCreate:     []string{"test1.txt", "test2.txt", "test3.txt"},
		},
		{
			Name:              "create 3 jpg files, expect to get 3 file names",
			ExpectedFileNames: []string{"test1.jpg", "test2.jpg", "test3.jpg"},
			FilesToCreate:     []string{"test1.jpg", "test2.jpg", "test3.jpg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := os.MkdirAll(TestFilesPath, os.ModePerm)
			if err != nil {
				t.Fatalf("Error creating directory: %v\n", err)
			}

			for _, fileName := range tt.FilesToCreate {
				filePath := filepath.Join(TestFilesPath, fileName)
				file, err := os.Create(filePath)
				if err != nil {
					t.Fatalf("Error creating file: %v\n", err)
				}
				defer file.Close()
			}

			got, err := GetFilesInADirectory(TestFilesPath)
			if err != nil {
				t.Fatalf("Error getting file names: %v", err)
			}

			if !reflect.DeepEqual(got, tt.ExpectedFileNames) {
				t.Fatalf("expected file names %+v, got %+v", tt.ExpectedFileNames, got)
			}

			os.RemoveAll(TestFilesPath)

		})
	}
}

func TestCreateArchive(t *testing.T) {
	tests := []struct {
		Name              string
		ExpectedFileNames []string
		FilesToCreate     []string
	}{
		{
			Name:              "create 3 txt files, expect to get 3 file names",
			ExpectedFileNames: []string{"test1.txt", "test2.txt", "test3.txt"},
			FilesToCreate:     []string{"test1.txt", "test2.txt", "test3.txt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := os.MkdirAll(TestFilesPath, os.ModePerm)
			if err != nil {
				t.Fatalf("Error creating directory: %v\n", err)
			}

			for _, fileName := range tt.FilesToCreate {
				filePath := filepath.Join(TestFilesPath, fileName)
				file, err := os.Create(filePath)
				if err != nil {
					t.Fatalf("Error creating file: %v\n", err)
				}
				defer file.Close()
			}

			err = CreateAnArchive(TestFilesPath)
			if err != nil {
				t.Fatalf("coudlnt create an archive with error %v", err)
			}
		})
	}
}
