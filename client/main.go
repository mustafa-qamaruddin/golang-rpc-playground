package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

const serverAddress string = "127.0.0.1"

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	AssignedBy  string `json:"assignedBy"`
}

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	todo := Todo{}
	client.Call("TodoService.Create", &Todo{
		ID:          "123",
		Title:       "Implement Go RPC",
		Description: "A user could create a Todo Model",
		Assignee:    "abc",
		AssignedBy:  "xyz",
	}, &todo)
	log.Println(fmt.Sprintf("Response: [%s] %s %s %s", time.Now(), todo.ID, todo.Title, todo.Description))
}
