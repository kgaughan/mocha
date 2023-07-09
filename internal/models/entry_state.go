package models

import "database/sql"

type EntryState struct {
	Entry  Entry
	User   User
	ReadAt sql.NullTime
}
