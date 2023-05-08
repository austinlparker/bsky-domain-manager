package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
	dropTable()
	createTable()
}
func dropTable() {
	sqlStmt := `DROP TABLE IF EXISTS kv;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table dropped")
}

func createTable() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS kv (
		id SERIAL PRIMARY KEY,
		handle TEXT NOT NULL,
		key TEXT NOT NULL,
		value TEXT NOT NULL,
		UNIQUE (handle, key)
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created")
}

func dbGet(handle string, key string) (string, error) {
	var value string
	err := db.QueryRow("SELECT value FROM kv WHERE handle=$1 AND key=$2", handle, key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no value found for handle=%s, key=%s", handle, key)
		}
		return "", err
	}
	return value, nil
}

func dbGetKeyForHandle(handle string) (string, error) {
	var value string
	err := db.QueryRow("SELECT value FROM kv WHERE handle=$1", handle).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no value for for handle=%s", handle)
		}
		return "", err
	}
	return value, nil
}

func dbSet(handle string, key string, value string) error {
	_, err := db.Exec("INSERT INTO kv (handle, key, value) VALUES ($1, $2, $3) ON CONFLICT (handle, key) DO UPDATE SET value = $3", handle, key, value)
	if err != nil {
		return err
	}
	return nil
}
