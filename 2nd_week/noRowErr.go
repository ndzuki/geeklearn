package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// define QueryError struct as error interface
type QueryError struct {
	Query string
	Err   error
}

// define Unwrap functionily return unwraped error
func (e *QueryError) Unwrap() error { return e.Err }

type production struct {
	id    int
	name  string
	price float64
}

func main() {
	var product production
	var qerr *QueryError

	db, err := sql.Open("mysql", "root1:admin@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal("Unable to Connect DB .", err)
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM prod")
	if err != nil {
		if errors.Is(err, &qerr.Err) {
			log.Fatal("Error when fetching prod table", qerr.Unwrap())
		} else {
			fmt.Println(err)
		}
	}

	defer results.Close()
	for results.Next() {

		err = results.Scan(&product.id, &product.name, &product.price)
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", product.id, product.name, product.price)
	}

	err = db.QueryRow("SELECT * FROM prod WHERE id = 3").Scan(&product.id, &product.name, &product.price)
	if err != nil {
		log.Fatal("Unable to parse row:", err)
	}
	fmt.Printf("ID: %d\nName: %s\nPrice: %0.2f\n", product.id, product.name, product.price)
}
