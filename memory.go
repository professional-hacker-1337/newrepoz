package main

import "fmt"

var tasks = make(map[int]*Task)

func Add(task *Task) error {

	id := len(tasks) + 1 // вычисляем айди(по количеству задач)
	task.ID = id         // перезаписываем айди в текущую мапу
	tasks[id] = task     // добавляем задачу в новую мапу по новому айди
	return nil
}

func List(status string) ([]*Task, error) {

	for _, task := range tasks {
		if status == "" ||
			(status == "todo" && task.Status == "TODO") ||
			(status == "done" && task.Status == "DONE") {
			fmt.Println("[", task.Status, "]", task.Title, "(", task.CreatedAt.Format("2006-01-02"), ")")
		}
	}

	return []*Task{}, nil

}

func Get(id int) (*Task, error) {
	task, exists := tasks[id]
	if !exists {
		return nil, fmt.Errorf("задача с ID %d не найдена", id)
	}

	fmt.Println(task.ID, "[", task.Status, "]", task.Title, "(", task.CreatedAt.Format("2006-01-02"), ")")
	return task, nil
}

func Update(task *Task) error {

	var newTitle string
	fmt.Scan(&newTitle)
	task.Title = newTitle
	return nil
}

func Delete(id int) error {

	delete(tasks, id)
	return nil

}
