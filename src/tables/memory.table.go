package tables

import (
	"database/sql"
	"fmt"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/database"
	"noraclock/v2/src/exception"
	"noraclock/v2/src/models"
)

type memoryTable struct{}

// Memory : Struct that encapsulates all table operations on Memory.
var Memory = &memoryTable{}

func (m *memoryTable) GetByID(id string) (*models.Memory, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, configs.Postgres.MemoryTableName)

	memory := &models.Memory{}
	err := database.GetPostgreSQL().QueryRow(query, id).Scan(
		&memory.ID,
		&memory.Title,
		&memory.Body,
		&memory.CreatedAt,
		&memory.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.MemoryNotFound("")
		}
		return nil, err
	}

	return memory, nil
}

func (m *memoryTable) Insert(id string, title string, body string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s ("id","title","body") VALUES ($1,$2,$3)`,
		configs.Postgres.MemoryTableName,
	)

	_, err := database.GetPostgreSQL().Exec(query, id, title, body)
	return err
}
