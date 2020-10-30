package main

import (
	"note/initialize"
)

func main() {
	db := initialize.InitPostgres()
	initialize.RegisterModel(db)
	defer db.Close()
}
