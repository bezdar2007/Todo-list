package main

import "fmt"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task

func AddTask(title string) Task {
	newID := len(tasks) + 1
	newTask := Task{ID: newID, Title: title}
	tasks = append(tasks, newTask)
	fmt.Printf("Добавлена задача: %v\n", newTask)
	return newTask
}

func GetAllTasks() []Task {
	return tasks
}

func DeleteTask(id int) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
