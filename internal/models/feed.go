package models

import (
	"database/sql"
)

type Feed struct {
	ID           uint `gorm:"primaryKey"`
	UUID         string
	Title        string
	FeedURL      string
	SiteURL      string
	CheckedAt    sql.NullTime
	ETag         sql.NullString
	LastModified sql.NullTime
}
