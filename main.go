package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/YukiBobier/practice-sqlc/internal/classicmodels"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/classicmodels")
	if err != nil {
		log.Fatal(err)
	}

	queries := classicmodels.New(db)

	ctx := context.Background()
	customer, err := queries.GetCustomer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}
