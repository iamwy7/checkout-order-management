package main

import "github.com/iamwy7/meli-challenge/orders/adapters/inputs/cli"

//	@title			Orders API
//	@version		1.0
//	@description	This is a server for managing orders.
//	@host			localhost:8080
//	@BasePath		/
func main() {
	cli.Execute()
}
