package product

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productStore struct {
	db *mongo.Database
}

const productCollection = "products"

func NewProductStore(db *mongo.Database) *productStore {
	return &productStore{db: db}
}

func (store *productStore) NewProduct(product *Product) {
	store.db.Collection(productCollection).InsertOne(context.TODO(), product)
}

func (store *productStore) GetAllProduct() []Product {
	cur, _ := store.db.Collection(productCollection).Find(context.TODO(), bson.D{{}})
	products := []Product{}

	for cur.Next(context.TODO()) {
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {
			return nil
		}

		products = append(products, elem)
	}

	return products
}
