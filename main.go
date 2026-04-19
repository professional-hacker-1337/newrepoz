package main

import "fmt"

func main() {

	for {
		var choice string
		var param string

		fmt.Scan(&choice, &param)

		if choice == "add" {

			task := Task{
				Title:  param,
				Status: "TODO", // "TODO" или "DONE"
			}

			Add(&task)

		}
		if choice == "list" {
			List(param)
		}

		if choice == "get" {
			var id int
			fmt.Scan(&id) // Считываем ID от пользователя
			Get(id)
		}
		if choice == "edit" {
			var id int
			fmt.Scan(&id)

			task, exists := tasks[id]
			if exists {
				Update(task) // передаём task в функцию
			}
		}

		if choice == "delete" {
			var id int
			fmt.Scan(&id)
			Delete(id)
		}
		if choice == "exit" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
