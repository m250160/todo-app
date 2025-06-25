package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	panic("unimplemented")
}

func ListTasks() {
	panic("unimplemented")
}

func CompleteTask(id int) {
	var tasks []Task
	var err error
	var i int
	tasks, err = loadTasks()
	if err != nil {
		fmt.Println("Error!")
		return
	}
	for i = 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Done = true
			break
		}
	}
	if i == len(tasks) {
		fmt.Println("No ID!")
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
