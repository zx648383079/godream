package models

import "time"

type BookLog struct {
	ID        uint
	BookID    int
	Hits      int
	CreatedAt time.Time
}
