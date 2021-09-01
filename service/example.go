package service

import (
	"log"
	"nutrition/database"
	"nutrition/database/mongo"
	inf "nutrition/infrastructure"
)

type ExampleService interface {
}

type exampleServiceIml struct {
	exampleRepository database.ExampleRepository
	InfoLog           *log.Logger
	ErrorLog          *log.Logger
}

func NewExampleService() ExampleService {
	dbStore := inf.NewMongoDatastore()
	collection := ""
	return &exampleServiceIml{
		exampleRepository: mongo.NewExampleRepository(collection, dbStore),
		InfoLog:           inf.InfoLog,
		ErrorLog:          inf.ErrLog,
	}
}
