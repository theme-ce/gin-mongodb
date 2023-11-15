package store

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/theme-ce/gin-mongodb/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBStore struct {
	collection *mongo.Collection
}

type MongoDBStorer interface {
	product.ProductHandler
}

func NewMongoDBStore(collection *mongo.Collection) MongoDBStorer {
	return &mongoDBStore{collection: collection}
}

func (mongo *mongoDBStore) InsertProduct(ctx *gin.Context) {
	productName := ctx.Param("name")

	_, err := mongo.collection.InsertOne(context.TODO(), product.Product{
		ProductName: productName,
	})
	if err != nil {
		panic(err)
	}
}
