package main

import (
	"api"
	"fmt"
	"net/http"
)

func main() {
	newRouter := api.NewRouter()
	err := http.ListenAndServe(":8080", newRouter)
	if err != nil {
		fmt.Printf("encountered error while listeningon port 8080: %s", err.Error())
	}
	fmt.Println("listening on port 8080")
}
