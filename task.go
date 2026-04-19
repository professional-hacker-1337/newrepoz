package main

import "time"

type Task struct {
	ID        int
	Title     string
	Status    string // "TODO" или "DONE"
	CreatedAt time.Time
}
