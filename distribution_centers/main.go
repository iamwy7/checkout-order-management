package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"                             // Database driver
	_ "github.com/iamwy7/meli-challenge/distribution_centers/docs" // Import generated Swagger docs
	api_handler "github.com/iamwy7/meli-challenge/distribution_centers/handler"
	"github.com/iamwy7/meli-challenge/distribution_centers/repo"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Distribution Centers API
// @version		1.0
// @description	This is a server for managing disctribution centers.
// @host			localhost:8080
// @BasePath		/
func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	apiPort := os.Getenv("API_PORT")
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" || apiPort == "" {
		log.Fatal("one or more missing env vars, please check application documentation")
	}
	dbConnection, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()
	dcRepo, err := repo.NewDistributionCenterDBFactory(dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	dcHandler := api_handler.NewDistributionCentersHandler(dcRepo)

	// Create the multiplex router
	r := http.NewServeMux()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("UP"))
	})
	r.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	r.HandleFunc("/distributioncenters", dcHandler.GetDistributionCenter)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", apiPort),
		Handler: r,
	}

	// Channel to listen for OS signals
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Shutdown the server gracefully
		log.Println("starting gracefully shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("error at graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Starting
	log.Printf("starting server on port %v \n", apiPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("error to start the HTTP server: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("server shutted down, cya :)")

}
