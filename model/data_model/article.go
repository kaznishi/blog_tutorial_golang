package data_model

import "time"

type Article struct {
	ID int
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}
