package main

import (
	"togo/cmd"
	"togo/internal/storage"
	"togo/internal/task/service"
)

func main() {
	store := &storage.JSONStorage{File: "tasks.json"}
	service.InitStore(store)
	cmd.Execute()
}
