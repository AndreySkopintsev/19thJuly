package api

import (
	"common"
	"encoding/json"
	"net/http"
)

func AddLinksToTask(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	taskId := queryParams.Get("taskid")

	newRequestBody := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&newRequestBody)
	if err != nil {
		WriteResponse(FailedTaskId, false, err.Error(), w, http.StatusBadRequest)
		return
	}

	if task, ok := createdTasks[taskId]; !ok {
		WriteResponse(FailedTaskId, false, ErrNoTaskFound.Error(), w, http.StatusBadRequest)
		return
	} else {
		if len(task) >= common.NumberOfLinks {
			WriteResponse(FailedTaskId, false, ErrTaskLinks.Error(), w, http.StatusBadRequest)
			return
		}
		for _, link := range newRequestBody.Links {
			task = append(task, link)
			if len(task) == common.NumberOfLinks {
				break
			}
		}
		createdTasks[taskId] = task
		WriteResponse(taskId, true, LinkAdded, w, http.StatusOK)
	}
}
