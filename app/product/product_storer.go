package product

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type productStore struct {
	collection *mongo.Collection
}

func NewProductStore(collection *mongo.Collection) *productStore {
	return &productStore{collection: collection}
}

func (store *productStore) NewProduct(product *Product) {
	_, err := store.collection.InsertOne(context.TODO(), product)
	if err != nil {
		panic(err)
	}
}
