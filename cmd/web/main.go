// git remote add origin https://github.com/ub2013210024/hello.git
// git branch -M main
// git push -u origin main
// A handler function is a that executes when activted by a route

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Crate a new type
type application struct{}

func main() {
	// Create a flag for specifying the port number for when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("LEMON_DB_DSN"), "PostgreSQL dsn")
	flag.Parse()
	// Create an instance of the application type
	app := &application{}
	// Create an instance of the connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	log.Println("database connection pool established")

	// Create a customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
