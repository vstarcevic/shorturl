package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shorturl/api"
	"shorturl/database"
)

func main() {

	dsn := os.Getenv("DSN")
	baseUrl := os.Getenv("BASE_URL")

	dbConn := database.ConnectToDB(dsn)
	defer dbConn.Close()

	cfg := api.Config{
		Db:      dbConn,
		BaseUrl: baseUrl,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "9000"),
		Handler: api.Routes(&cfg),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
