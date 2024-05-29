package main

import (
	"demo08/database"
	"demo08/router"
)

func main() {
	database.Init()
	r := router.Init()
	r.Run(":8080")
}
