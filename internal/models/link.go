package models

import "database/sql"

type Link struct {
	ID       uint `gorm:"primaryKey"`
	Entry    Entry
	Ref      string // alternate, related, enclosure, via
	Title    string
	URL      string
	Size     uint
	MIMEType sql.NullString
}
