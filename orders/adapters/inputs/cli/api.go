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
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"               // Database driver
	_ "github.com/iamwy7/meli-challenge/orders/docs" // Import generated Swagger docs
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/iamwy7/meli-challenge/orders/adapters/inputs/api_handlers"
	"github.com/iamwy7/meli-challenge/orders/adapters/outputs/db_repository"
	"github.com/iamwy7/meli-challenge/orders/adapters/outputs/http_client"
	"github.com/iamwy7/meli-challenge/orders/application/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port string

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Api is the way to interact with application.",
	Long:  `Api command enable the '/orders' endpoint, that can receive a json to create an order or get one by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbHost := viper.GetString("DB_HOST")
		dbPort := viper.GetString("DB_PORT")
		dbUser := viper.GetString("DB_USER")
		dbPass := viper.GetString("DB_PASSWORD")
		dbName := viper.GetString("DB_NAME")
		if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
			log.Fatal("one or more missing env vars, please check application documentation")
		}
		dbConnection, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPass, dbHost, dbPort, dbName))
		if err != nil {
			log.Fatal(err)
		}
		defer dbConnection.Close()

		// Create the repositories
		dcUrl := viper.GetString("DC_URL")
		dcPort := viper.GetString("DC_PORT")
		if dcUrl == "" || dcPort == "" {
			log.Fatal("one or more missing env vars, please check application documentation")
		}
		dcRepo := http_client.NewDistributionCenterAdapterFactory(fmt.Sprintf("http://%v:%v", dcUrl, dcPort))
		ordeRepo, err := db_repository.NewMySqlOrderAdapterFactory(dbConnection)
		if err != nil {
			log.Fatal(err)
		}

		// Create the use cases
		createOrderUseCase := usecase.NewCreateOrderUseCase(ordeRepo, dcRepo)
		getOrderUseCase := usecase.NewGetOrderUseCase(ordeRepo)

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

	viper.AutomaticEnv() // Automatically read environment variables

	// Define environment variable bindings
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DC_URL")
	viper.BindEnv("DC_PORT")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	apiCmd.Flags().StringVarP(&port, "port", "p", ":8080", "Change port to be used on api")
}
