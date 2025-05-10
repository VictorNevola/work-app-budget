package main

import "github.com/VictorNevola/work-app-budget/configuration/database"

func main() {
	_, disconnectDB := database.NewMongoDB()

	defer func() {
		disconnectDB()
	}()
}
