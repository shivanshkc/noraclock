package tables

import (
	"fmt"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/database"
)

type memoryTable struct{}

// Memory : Struct that encapsulates all table operations on Memory.
var Memory = &memoryTable{}

func (m *memoryTable) Insert(id string, title string, body string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s ("id","title","body") VALUES ($1,$2,$3)`,
		configs.Postgres.MemoryTableName,
	)

	_, err := database.GetPostgreSQL().Exec(query, id, title, body)
	return err
}
