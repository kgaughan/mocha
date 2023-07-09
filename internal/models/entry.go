package models

import (
	"database/sql"
	"time"
)

type Entry struct {
	ID          uint `gorm:"primaryKey"`
	Feed        Feed
	UUID        string
	URL         string
	Title       string
	Summary     sql.NullString
	Content     sql.NullString
	Author      sql.NullString
	PublishedAt time.Time
	UpdatedAt   time.Time
}
