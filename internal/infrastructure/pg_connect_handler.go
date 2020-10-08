package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"

	_ "github.com/lib/pq"
)

// PostgresConnector is shared connection to DB
type PostgresConnector struct {
	Conn *sql.DB
	Pool *pgxpool.Pool
}

// SQLRows results set
type SQLRows struct {
	Rows *sql.Rows
}

type DBConnector interface {
	Execute(statement string, args  ...interface{}) (sql.Result, error)
	Query(statement string) (Rows, error)
	QueryRow(statement string, args  ...interface{}) pgx.Row
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
}

type Result interface {
	sql.Result
}

func (pgConn *PostgresConnector) QueryRow(statement string, args ...interface{}) pgx.Row {
	conn, err := pgConn.Pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error acquiring connection:", err)
		os.Exit(1)
	}
	defer conn.Release()
	return conn.QueryRow(context.Background(), statement, args...)
}

func (pgConn *PostgresConnector) Execute(statement string, args ...interface{}) (sql.Result, error) {
	return pgConn.Conn.Exec(statement, args...)
}

func (pgConn *PostgresConnector) Query(statement string) (Rows, error) {
	rows, err := pgConn.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(SQLRows), err
	}

	rowsSet := new(SQLRows)
	rowsSet.Rows = rows

	return rowsSet, nil
}

func (r SQLRows) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SQLRows) Next() bool {
	return r.Rows.Next()
}

