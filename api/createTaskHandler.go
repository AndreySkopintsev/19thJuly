package api

import (
	"common"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const (
	FailedTaskId  = "-1"
	TaskAddedText = "Task created"
	LinkAdded     = "Link added"
)

var (
	ErrBadJson     = errors.New("couldnt parse received JSON, please check and send again")
	ErrNumOfLinks  = errors.New(fmt.Sprintf("too many links provided, please reduce the number of links to %d", common.NumberOfLinks))
	ErrNumOfTasks  = errors.New(fmt.Sprintf("there are already %d tasks in progress, please try adding new tasks later", common.NumberOfTasks))
	ErrNoTaskFound = errors.New("no task with provided id was found")
	ErrTaskLinks   = errors.New(fmt.Sprintf("this task already has %d links", common.NumberOfLinks))
)

var createdTasks map[string][]string = map[string][]string{}

type RequestBody struct {
	Links []string
}

type ResponseBody struct {
	Success bool
	TaskId  string
	Message string
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if len(createdTasks) == 3 {
		WriteResponse(FailedTaskId, false, ErrNumOfTasks.Error(), w, http.StatusBadRequest)
		return
	}
	newRequestBody := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&newRequestBody)
	if err != nil {
		WriteResponse(FailedTaskId, false, err.Error(), w, http.StatusBadRequest)
		return
	}

	if len(newRequestBody.Links) > common.NumberOfLinks {
		WriteResponse(FailedTaskId, false, ErrNumOfLinks.Error(), w, http.StatusBadRequest)
		return
	}
	newTaskId := uuid.New().String()

	createdTasks[newTaskId] = newRequestBody.Links
	WriteResponse(newTaskId, true, TaskAddedText, w, http.StatusOK)
}

func WriteResponse(taskId string, successStatus bool, message string, w http.ResponseWriter, httpStatus int) {
	newResp := ResponseBody{
		Success: successStatus,
		TaskId:  taskId,
		Message: message,
	}
	respBBody, _ := json.Marshal(newResp)
	w.WriteHeader(httpStatus)
	w.Write(respBBody)
}
