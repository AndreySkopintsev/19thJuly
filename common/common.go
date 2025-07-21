package common

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

var NumberOfTasks int
var NumberOfLinks int

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	inNumOfLinks, err := strconv.Atoi(os.Getenv("NUMBER_OF_LINKS"))
	if err != nil {
		fmt.Println("couldnt get number of links from the environment")
	}
	inNumOfTasks, err := strconv.Atoi(os.Getenv("NUMBER_OF_TASKS"))
	if err != nil {
		fmt.Println("couldnt get number of links from the environment")
	}
	NumberOfLinks = inNumOfLinks
	NumberOfTasks = inNumOfTasks
}
