package main

import (
	"fmt"

	drivers "github.com/ryoichi-k/go_todo/drivers"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world!!!!!!!!")
	_, err := drivers.NewMysqlDriver()
	if err != nil {
		fmt.Println("n")
	}
}
