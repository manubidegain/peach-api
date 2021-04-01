package io

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"peach-core/entities"

	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func OpenMySQL() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No env file provided")
	}
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:6612)/%s", user, password, database))
	if err != nil {
		panic(err.Error())
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []entities.Product
	result, err := db.Query("SELECT * from products")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var product entities.Product
		err := result.Scan(&product.ProductID, &product.Provider, &product.Link, &product.Brand, &product.Name, &product.Category, &product.Stock, &product.Price)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO products(provider,link,brand,name,category,stock,price) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var product entities.Product
	json.Unmarshal(body, &product)
	_, err = stmt.Exec(product.Provider, product.Link, product.Brand, product.Name, product.Category, product.Stock, product.Price)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New product was created")
}
