package main

import (
	"togo/cmd"
	"togo/internal/task"
)

func main() {
	task.LoadFromFile()
	cmd.Execute()
}
