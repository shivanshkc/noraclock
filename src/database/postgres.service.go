package database

import (
	ctx "context"
	"database/sql"
	"fmt"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/logger"
	"time"
)

var log, _ = logger.Get()

// PostgreSQL : Database connection struct.
type PostgreSQL struct {
	DB *sql.DB
}

var pg = &PostgreSQL{}

// ConnectPostgreSQL : Establishes or refreshes connection with PostgreSQL.
func ConnectPostgreSQL() (*sql.DB, error) {
	connectionInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		configs.Postgres.Host,
		configs.Postgres.Port,
		configs.Postgres.User,
		configs.Postgres.Password,
		configs.Postgres.DBName,
	)

	db, err := sql.Open("pgx", connectionInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Sugar().Infof("Connected to PostgreSQL at %s:%s", configs.Postgres.Host, configs.Postgres.Port)
	pg.DB = db

	if err = pg.initiate(); err != nil {
		log.Sugar().Infof("Failed to complete Post-Connect operations on PostgreSQL. %s", err.Error())
		return nil, err
	}

	return db, nil
}

// GetPostgreSQL : Get database struct.
func GetPostgreSQL() *PostgreSQL {
	return pg
}

// Exec : Executes a query without returning rows. Returns the number of affected rows.
func (p *PostgreSQL) Exec(query string, args ...interface{}) (int64, error) {
	context, cancelFunc := getTimeoutContext()
	defer cancelFunc()

	result, err := p.DB.ExecContext(context, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// Query : Executes a query and returns rows.
func (p *PostgreSQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	context, cancelFunc := getTimeoutContext()
	defer cancelFunc()

	return p.DB.QueryContext(context, query, args...)
}

// QueryRow : Executes a query and returns a single row.
func (p *PostgreSQL) QueryRow(query string, args ...interface{}) *sql.Row {
	context, cancelFunc := getTimeoutContext()
	defer cancelFunc()

	return p.DB.QueryRowContext(context, query, args...)
}

func (p *PostgreSQL) initiate() error {
	if _, err := p.Exec(createUpdatedAtFunc()); err != nil {
		return err
	}
	if _, err := p.Exec(createMemoryTableIfNotExists()); err != nil {
		return err
	}
	if _, err := p.Exec(createUpdatedAtTrigger()); err != nil {
		return err
	}
	return nil
}

func getTimeoutContext() (ctx.Context, ctx.CancelFunc) {
	return ctx.WithTimeout(
		ctx.Background(),
		time.Duration(configs.Postgres.RequestTimeout)*time.Second,
	)
}
