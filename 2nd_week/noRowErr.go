package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

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
		var (
			id    int
			name  string
			price float64
		)
		err = results.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", id, name, price)
	}

	name, id := "Laptop", 3
	err = db.QueryRow("SELECT * FROM prod WHERE name = ? and id = ?", name, id).Scan(&id, &name)
	if err != nil {
		log.Printf("product %s (id=%d) not found. \n%w", name, id, err)
	}
}
