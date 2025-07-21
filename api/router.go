package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Path    string
	Action  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// TODO add handlers
var Routes []Route = []Route{
	{Name: "create task", Path: "/createTask", Action: "POST", Handler: CreateTaskHandler},
	{Name: "add link to task", Path: "/addLink", Action: "POST", Handler: func(w http.ResponseWriter, r *http.Request) {}},
	{Name: "get task status", Path: "/getTaskStatus", Action: "GET", Handler: func(w http.ResponseWriter, r *http.Request) {}},
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Action)
	}
	return r
}
