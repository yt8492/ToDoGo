package main

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type TodoHandler struct {
	todoMap map[string]string
}

func (handler *TodoHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		for id, todo := range handler.todoMap {
			fmt.Fprintf(writer, "%s: %s\n", id, todo)
		}
	case "POST":
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(request.Body)
		body := bufBody.String()
		id := uuid.New().String()
		handler.todoMap[id] = body
		fmt.Fprintf(writer, id)
	case "DELETE":
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(request.Body)
		id := bufBody.String()
		delete(handler.todoMap, id)
		fmt.Fprintf(writer, "%s deleted", id)
	}
}

func main() {
	todoHandler := TodoHandler{
		todoMap: map[string]string{},
	}
	http.HandleFunc("/todo", todoHandler.Handle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
