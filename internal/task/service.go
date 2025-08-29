package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var tasks []Task
var lastID int

const fileName = "tasks.json"

func Create(title, description string) Task {
	lastID++
	t := Task{
		ID:          lastID,
		Title:       title,
		Description: description,
		Status:      "new",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	tasks = append(tasks, t)
	saveToFile()
	return t
}

func List() []Task {
	return tasks
}

func Update(id int, title, description, status string) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			if title != "" {
				tasks[i].Title = title
			}
			if description != "" {
				tasks[i].Description = description
			}
			if status != "" {
				tasks[i].Status = status
			}
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			saveToFile()
			return true
		}
	}
	return false
}

func Delete(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveToFile()
			return true
		}
	}
	return false
}

func Find(id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func saveToFile() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации JSON:", err)
		return
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		fmt.Println("Ошибка записи файла:", err)
	}
}

func LoadFromFile() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return // файла нет — начинаем с пустого списка
		}
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}

	for _, t := range tasks {
		if t.ID > lastID {
			lastID = t.ID
		}
	}
}
