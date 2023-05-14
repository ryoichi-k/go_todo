package drivers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

// builder
func NewMysqlDriver() (*Mysql, error) {
	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/go_todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("connection test failed")
		return nil, err
	}
	fmt.Println("connection test success")
	m := &Mysql{db: db}
	return m, nil
}

func (m *Mysql) Query(query string) (*sql.Rows, error) {
	rows, err := m.db.Query(query)
	if err != nil {
		log.Fatalf("Query db.Query error err:%v", err)
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}

func (m *Mysql) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	rows, err := m.db.QueryContext(ctx, query, args)
	if err != nil {
		log.Fatalf("Query db.QueryContext error err:%v", err)
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}

type Rows struct {
	Rows *sql.Rows
}

func (r *Rows) Close() {
	defer r.Rows.Close()
}

func (r *Rows) Next() {
	r.Rows.Next()
}

func (r *Rows) Scan(dest ...any) error {
	err := r.Rows.Scan(dest)
	if err != nil {
		return err
	}
	return nil
}
