package main

import (
	"github.com/VictorNevola/work-app-budget/api/http"
	"github.com/VictorNevola/work-app-budget/internal/adapters/db/mongo"
	"github.com/go-playground/validator/v10"
)

func main() {
	mongoConnection, disconnectDB := mongo.NewMongoDB()
	validator := validator.New()

	http.NewFiberServer(mongoConnection, validator)

	defer func() {
		disconnectDB()
	}()
}
