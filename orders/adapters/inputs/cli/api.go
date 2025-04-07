/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"                        // Database driver
	_ "github.com/iamwy7/meli-challenge/orders/adapters/docs" // Import generated swagger docs
	"github.com/iamwy7/meli-challenge/orders/adapters/inputs/api_handlers"
	"github.com/iamwy7/meli-challenge/orders/adapters/outputs/db_repository"
	"github.com/iamwy7/meli-challenge/orders/adapters/outputs/http_client"
	"github.com/iamwy7/meli-challenge/orders/application/usecase"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
)

var port string

// @title Orders API
// @version 1.0
// @description This is a server for managing orders.
// @host localhost:8080
// @BasePath /
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Api is the way to interact with application.",
	Long:  `Api command enable the '/orders' endpoint, that can receive a json to create an order or get one by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create the database connection
		db_conn, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/orders_db")
		if err != nil {
			log.Fatal(err)
		}
		defer db_conn.Close()

		// Create the repositories
		dc_repo := http_client.NewDistributionCenterAdapterFactory("http://wiremock:8081")
		order_repo, err := db_repository.NewMySqlOrderAdapterFactory(db_conn)
		if err != nil {
			log.Fatal(err)
		}

		// Create the use cases
		createOrderUseCase := usecase.NewCreateOrderUseCase(order_repo, dc_repo)
		getOrderUseCase := usecase.NewGetOrderUseCase(order_repo)

		// Create the handlers
		ordersHandler := api_handlers.NewOrderHandler(
			createOrderUseCase,
			getOrderUseCase,
		)

		// Create the multiplex router
		r := http.NewServeMux()
		r.HandleFunc("/swagger/", httpSwagger.WrapHandler)
		r.HandleFunc("/orders/{orderId}", ordersHandler.GetOrder)
		r.HandleFunc("POST /orders", ordersHandler.CreateOrder)

		server := &http.Server{
			Addr:    port,
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
		log.Printf("starting server on port %v \n", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("error to start the HTTP server: %v\n", err)
		}

		<-idleConnsClosed
		log.Println("server shutted down, cya :)")
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	apiCmd.Flags().StringVarP(&port, "port", "p", ":8080", "Change port to be used on api")
}
