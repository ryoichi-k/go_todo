package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryoichi-k/go_todo/algo"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(db)/go_todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("connection test failed")
	}
	fmt.Println("connection test success")
	rows, err := db.Query("select name from users where id = ? ", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/age", handlerAge)
	// Web サーバーを起動する
	fmt.Println("server runnning...")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ナイス！君は神だよ！自信もって！")
}
func handlerAge(w http.ResponseWriter, r *http.Request) {
	var human = "yajuu"
	fmt.Printf("24歳です。%s \n", human)
	algo.MovementPattern()
}
