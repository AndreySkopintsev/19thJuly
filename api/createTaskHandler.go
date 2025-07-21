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
	LinkNotReady  = "Link isnt ready yet, the task is still in progress"
)

var (
	ErrBadJson            = errors.New("couldnt parse received JSON, please check and send again")
	ErrNumOfLinks         = errors.New(fmt.Sprintf("too many links provided, please reduce the number of links to %d", common.NumberOfLinks))
	ErrNumOfTasks         = errors.New(fmt.Sprintf("there are already %d tasks in progress, please try adding new tasks later", common.NumberOfTasks))
	ErrNoTaskFound        = errors.New("no task with provided id was found")
	ErrTaskLinks          = errors.New(fmt.Sprintf("this task already has %d links", common.NumberOfLinks))
	ErrCouldntArchive     = errors.New("coudlnt create archive")
	ErrForbiddenExtension = errors.New("this extension is not supported")
)

var createdTasks map[string]Task = map[string]Task{}

type Task struct {
	Links       []string
	ArchiveLink string
}

type RequestBody struct {
	Links []string
}

type ResponseBody struct {
	Success     bool
	TaskId      string
	Message     string
	ArchiveLink string
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if len(createdTasks) == common.NumberOfTasks {
		WriteResponse(FailedTaskId, false, ErrNumOfTasks.Error(), w, http.StatusBadRequest, LinkNotReady)
		return
	}
	newRequestBody := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&newRequestBody)
	if err != nil {
		WriteResponse(FailedTaskId, false, err.Error(), w, http.StatusBadRequest, LinkNotReady)
		return
	}

	if len(newRequestBody.Links) > common.NumberOfLinks {
		WriteResponse(FailedTaskId, false, ErrNumOfLinks.Error(), w, http.StatusBadRequest, LinkNotReady)
		return
	}
	newTaskId := uuid.New().String()

	createdTasks[newTaskId] = Task{
		Links:       newRequestBody.Links,
		ArchiveLink: LinkNotReady,
	}
	WriteResponse(newTaskId, true, TaskAddedText, w, http.StatusOK, LinkNotReady)
}

func WriteResponse(taskId string, successStatus bool, message string, w http.ResponseWriter, httpStatus int, archiveLink string) {
	newResp := ResponseBody{
		Success:     successStatus,
		TaskId:      taskId,
		Message:     message,
		ArchiveLink: archiveLink,
	}
	respBBody, _ := json.Marshal(newResp)
	w.WriteHeader(httpStatus)
	w.Write(respBBody)
}
