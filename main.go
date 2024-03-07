package main

import (
	"fmt"
	"main/database"
)

func main() {
	err := database.DBConnect(database.GetDotEnv(), "tskr_dev_db")
	if err != nil {
		fmt.Printf("Failed to connect to DB: %s", err)
		return
	}
	fmt.Println("Connected to DB")
}
