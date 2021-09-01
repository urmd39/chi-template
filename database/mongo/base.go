package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type baseRepository struct {
	primaryCollection *mongo.Collection
}
