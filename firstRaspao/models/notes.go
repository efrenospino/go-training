package models

import "time"

type Note struct {
	ID      int
	Content string
	Date    time.Time
}
