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

func (m *memoryTable) Get(limit int64, offset int64, skipBody bool) ([]*models.Memory, int, error) {
	bodySelect := `"body",`
	if skipBody {
		bodySelect = `'' as "body",`
	}

	query := fmt.Sprintf(
		`SELECT "id", "title", %s "createdAt", "updatedAt", COUNT(*) OVER() FROM %s ORDER BY "createdAt" DESC LIMIT $1 OFFSET $2`,
		bodySelect,
		configs.Postgres.MemoryTableName,
	)

	var memories []*models.Memory
	rows, err := database.GetPostgreSQL().Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer func() { _ = rows.Close() }()

	var count int
	for rows.Next() {
		memory := &models.Memory{}
		err := rows.Scan(
			&memory.ID,
			&memory.Title,
			&memory.Body,
			&memory.CreatedAt,
			&memory.UpdatedAt,
			&count,
		)

		if err != nil {
			return nil, 0, err
		}

		memories = append(memories, memory)
	}

	if memories == nil {
		memories = []*models.Memory{}
	}
	return memories, count, nil
}

func (m *memoryTable) Insert(id string, title string, body string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s ("id","title","body") VALUES ($1,$2,$3)`,
		configs.Postgres.MemoryTableName,
	)

	_, err := database.GetPostgreSQL().Exec(query, id, title, body)
	return err
}

func (m *memoryTable) UpdateByID(id string, updates map[string]interface{}) error {
	set := ""
	argPos := 1
	var values []interface{}

	v, exists := updates["title"]
	if exists {
		set += fmt.Sprintf(`"title" = $%d, `, argPos)
		argPos++
		values = append(values, v)
	}
	v, exists = updates["body"]
	if exists {
		set += fmt.Sprintf(`"body" = $%d, `, argPos)
		argPos++
		values = append(values, v)
	}

	// Removing the trailing comma and space from 'set'.
	set = set[:len(set)-2]
	values = append(values, id)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE "id" = $%d`, configs.Postgres.MemoryTableName, set, argPos)
	affected, err := database.GetPostgreSQL().Exec(query, values...)
	if err != nil {
		return err
	}

	if affected == 0 {
		return exception.MemoryNotFound("")
	}
	return nil
}

func (m *memoryTable) DeleteByID(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE "id" = $1`, configs.Postgres.MemoryTableName)

	affected, err := database.GetPostgreSQL().Exec(query, id)
	if err != nil {
		return err
	}

	if affected == 0 {
		return exception.MemoryNotFound("")
	}
	return nil
}
