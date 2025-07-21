package common

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lpernett/godotenv"
)

var NumberOfTasks int
var NumberOfLinks int
var AllowedExtensions []string

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	numOfLinks := os.Getenv("NUMBER_OF_LINKS")
	if numOfLinks == "" {
		NumberOfLinks = 3
	} else {
		inNumOfLinks, err := strconv.Atoi(numOfLinks)
		if err != nil {
			fmt.Println("couldnt get number of links from the environment")
		}
		NumberOfLinks = inNumOfLinks
	}

	numOfTasks := os.Getenv("NUMBER_OF_TASKS")
	if numOfTasks == "" {
		NumberOfTasks = 3
	} else {
		inNumOfTasks, err := strconv.Atoi(numOfTasks)
		if err != nil {
			fmt.Println("couldnt get number of links from the environment")
		}
		NumberOfTasks = inNumOfTasks
	}

	allowedExtensions := os.Getenv("ALLOWED_EXTENSIONS")
	if allowedExtensions != "" {
		extensionsSlice := strings.Split(allowedExtensions, ",")
		AllowedExtensions = extensionsSlice
	} else {
		AllowedExtensions = []string{".pdf", ".jpeg"}
	}

}
