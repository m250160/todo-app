package main

import (
	"encoding/json"
	"fmt"
	"os"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks,err := loadTasks()
	if err != nil {
		println("読み込みエラー:", err.Error())
	}
	var test Task
	test.ID = nextID(tasks)
	test.Title = title
	test.Done = false

	tasks = append(tasks, test)
	saveTasks(tasks)
}

func ListTasks() {
	tasks,err := loadTasks()
	if err != nil {
		fmt.Println("読み込みエラー:", err.Error())
	}
	for i := 0; i < len(tasks); i++ {
		var x string
		if tasks[i].Done{
			x = "[x]"
		} else {
			x = "[ ]"
		}
		fmt.Println(tasks[i].ID, ":", tasks[i].Title, x)
	}
}

func CompleteTask(id int) {
	var tasks []Task
	var err error
	tasks, err = loadTasks()
	if err != nil {
		fmt.Println("Error!")
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Done = true
		}
	}
	saveTasks(tasks)
}

func DeleteTask(id int) {
	var tasks []Task
	var err error
	tasks, err = loadTasks()
	if err != nil {
		fmt.Println("Error!")
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	saveTasks(tasks)
}
