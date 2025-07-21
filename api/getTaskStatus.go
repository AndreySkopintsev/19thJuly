package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	taskId := queryParams.Get("taskid")

	task, ok := createdTasks[taskId]
	if !ok {
		WriteResponse(FailedTaskId, false, ErrNoTaskFound.Error(), w, http.StatusBadRequest, LinkNotReady)
		return
	}

	if len(task.Links) < 3 {
		WriteResponse(taskId, true, LinkNotReady, w, http.StatusOK, LinkNotReady)
		return
	}

	for idx, link := range task.Links {
		err := DownloadFile("file"+strconv.Itoa(idx), link, taskId)
		if err != nil {
			fmt.Printf("error downloading file: %s", err.Error())
			fmt.Println("coudlnt download file from provided link")
			continue
		}

	}

	err := CreateAnArchive(taskId)
	if err != nil {
		fmt.Printf("archive couldnt be created with error %s", err.Error())
		WriteResponse(taskId, false, ErrCouldntArchive.Error(), w, http.StatusOK, LinkNotReady)
		return
	}

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	WriteResponse(taskId, true, "Archive created", w, http.StatusOK, currentPath+"/"+taskId)
}
