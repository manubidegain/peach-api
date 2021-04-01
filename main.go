package main

import (
	"fmt"
	"net/http"

	i "peach-core/io"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	i.OpenMySQL()
	router := i.NewRouter()
	fmt.Println("Api listening at port 8080")
	http.ListenAndServe(":8000", router)
}
