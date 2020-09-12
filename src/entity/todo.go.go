package entity

import "time"

type ToDo struct {
	Title string
	Done  bool
	Limit time.Time
}
