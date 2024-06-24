package main

import "learn-go/internal/db"

func main() {
	db.InitDB()

	defer db.DB.Close()

}
