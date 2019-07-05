package main

import (
	"SQL/actType/db"
	"SQL/actType/routes"
)

func main() {
	database := db.ConnectToDb()
	routes.Init(database)
}
