package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
)

func TestCreateTaskHandler(t *testing.T) {
	tests := []struct {
		Name          string
		NumberOfTasks int
		Links         []string
		ExpectedResp  ResponseBody
	}{
		{
			Name:          "add just one task with no more than 3 links",
			NumberOfTasks: 1,
			Links:         []string{"link1", "link2", "link3"},
			ExpectedResp: ResponseBody{
				Success: true,
				Message: TaskAddedText,
			},
		},
		{
			Name:          "add just one task with more than 3 links",
			NumberOfTasks: 1,
			Links:         []string{"link1", "link2", "link3", "link4"},
			ExpectedResp: ResponseBody{
				Success: false,
				Message: ErrNumOfLinks.Error(),
			},
		},
		{
			Name:          "add just three tasks with no more than 3 links",
			NumberOfTasks: 3,
			Links:         []string{"link1", "link2", "link3"},
			ExpectedResp: ResponseBody{
				Success: true,
				Message: TaskAddedText,
			},
		},
		{
			Name:          "add more than three tasks with no more than 3 links",
			NumberOfTasks: 4,
			Links:         []string{"link1", "link2", "link3"},
			ExpectedResp: ResponseBody{
				Success: false,
				Message: ErrNumOfTasks.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			for i := 1; i <= tt.NumberOfTasks; i++ {
				reqbBytes, _ := json.Marshal(RequestBody{
					Links: tt.Links,
				})
				bodyReader := bytes.NewReader(reqbBytes)
				req := httptest.NewRequest("POST", "/createTask", bodyReader)

				w := httptest.NewRecorder()

				CreateTaskHandler(w, req)

				resp := w.Result()
				respbBytes, _ := io.ReadAll(resp.Body)
				newResp := ResponseBody{}
				json.Unmarshal(respbBytes, &newResp)
				if i == tt.NumberOfTasks {
					if newResp.Success != tt.ExpectedResp.Success || newResp.Message != tt.ExpectedResp.Message {
						t.Fatalf("expected response %+v, got response %+v", tt.ExpectedResp, newResp)
					}
				}

			}

			createdTasks = map[string][]string{}
		})
	}
}
