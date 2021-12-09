package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type production struct {
	id    int
	name  string
	price float64
}

func main() {
	var product production

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal("Unable to Connect DB .", err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM prod")
	if err != nil {
		log.Fatal("Error when fetching prod table", err)
	}
	defer results.Close()

	for results.Next() {
		err = results.Scan(&product.id, &product.name, &product.price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", product.id, product.name, product.price)
	}

	const name, id = "Laptop", 3
	err = db.QueryRow("SELECT * FROM prod WHERE name = ? and id = ?", name, id).Scan(&product.id, &product.name, &product.price)
	if err != nil {
		log.Fatalf("product %s (id=%d) not found.", name, id)
	}
	fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", product.id, product.name, product.price)
}
