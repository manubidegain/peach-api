package main

import (
	"fmt"
	"net/http"

	i "peach-core/io"
	u "peach-core/usecases"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	i.OpenMySQL()
	router := i.NewRouter()
	fmt.Println("Api listening at port 8080")
	u.ScrapeUnicom()
	http.ListenAndServe(":8000", router)

}
