package mongo

import (
	"nutrition/database"
	"nutrition/infrastructure"
)

type exampleRepositoryIml struct {
	baseRepository
}

func NewExampleRepository(primaryCollection string, store *infrastructure.MongoDBStore) database.ExampleRepository {
	return &exampleRepositoryIml{
		baseRepository: baseRepository{
			primaryCollection: store.Database.Collection(primaryCollection),
		},
	}
}
