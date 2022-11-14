package main

import (
	"fmt"
	"net/http"
)

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}

func handleUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ID:1, Name:张三")
		fmt.Fprintln(w, "ID:2, Name:李四")
		fmt.Fprintln(w, "ID:3, Name:王五")
		/*uJson, _ := json.Marshal(users)
		w.Write(uJson)*/
	default:
		w.WriteHeader(http.StatusNotFound)
		/*fmt.Fprintln(w, "not found")*/
		fmt.Fprint(w, "\"message\":\"not found\"")
	}
}

type User struct {
	ID   int
	Name string
}
