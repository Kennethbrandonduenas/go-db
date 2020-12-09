package storage

import (
	"database/sql"
	"sync"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:123456@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v",err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v",err)
		}

		fmt.Println("conectado a postgres")
	})
}

//Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}