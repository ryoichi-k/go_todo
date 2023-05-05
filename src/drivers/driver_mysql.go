package drivers

import (
	"database/sql"
	"log"
)

type mysql struct {
	db *sql.DB
}

func (d *mysql) NewMysqlDriver() (*mysql, error) {
	db, err := sql.Open("mysql", "admin:pass@tcp(0.0.0.0:3306)/go_todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
		return nil, err
	}
	m := &mysql{db: db}
	return m, nil
}
