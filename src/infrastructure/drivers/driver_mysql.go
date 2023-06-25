package drivers

import (
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
	db, err := sql.Open("mysql", "root:root@tcp(db)/go_todo")
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
