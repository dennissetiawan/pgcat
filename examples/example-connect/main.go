package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Define the PostgreSQL connection string.

	// connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
	// 	"postgres",
	// 	"password",
	// 	"localhost",
	// 	6432,
	// 	"order")

	connStringSpse := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"epns",
		"epns",
		"localhost",
		6432,
		"spse")

	// Connect to the database.
	// Test the database connection.
	// Execute a sample query.
	// query1 := "SELECT id FROM \"order\" LIMIT 10"
	// query2 := "SELECT peg_nama FROM \"pegawai\" LIMIT 10"
	// newFunction(connString, query1)
	// newFunction(connStringSpse, query2)

	db, err := sql.Open("postgres", connStringSpse)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// queryX, paramsX, err := sq.Select("tablename").From("pg_catalog.pg_tables").Prefix("/* shard_id: 0 */").Suffix("WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'").ToSql()

	query1 := "/* shard_id: 0 */ SELECT id FROM \"order\" LIMIT 10;"
	query2 := "/* shard_id: 1 */ SELECT peg_nama FROM \"pegawai\" LIMIT 10;"
	// queryY, paramsY, err := sq.Select("tablename").Prefix("/* shard_id: 1 */").From("pg_catalog.pg_tables").Suffix("WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'").ToSql()
	// newFunction(connString, query1)
	// newFunction(connStringSpse, query2)
	// fmt.Println(queryX, paramsX)
	rows, err := db.Query(query1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("orderID: %v\n", i)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows2, err := db.Query(query2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	for rows2.Next() {
		var i string
		if err := rows2.Scan(&i); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v\n", i)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func newFunction(connString string, query string) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v\n", i)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
