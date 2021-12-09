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

	id := 3
	err = db.QueryRow("SELECT * FROM prod WHERE id = ?", id).Scan(&product.id, &product.name, &product.price)
	if err != nil {
		// if errors.Is(err, errors.New("no rows in result set")) {
		log.Fatalf("product `id=%d` not found.", id)
	}
	fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", product.id, product.name, product.price)
}
