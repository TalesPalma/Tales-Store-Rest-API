package main

import (
	"fmt"

	"github.com/TalesPalma/gin-golang-rest/database"
	"github.com/TalesPalma/gin-golang-rest/routes"
)

func main() {
	database.ConnectionDatabase()
	routes.HandleRequests()
	fmt.Println("Bye ...")
}
