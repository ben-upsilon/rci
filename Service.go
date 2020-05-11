package main

import "net/http"

func CheckTaskStatus(w http.ResponseWriter, r *http.Request) {

}

type Task struct {
	id    int64
	name  string
	state string
}
