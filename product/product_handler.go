package product

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductHandler struct {
	client *mongo.Client
}

func NewProductHandler(client *mongo.Client) *ProductHandler {
	return &ProductHandler{client: client}
}

func (handler *ProductHandler) InsertProduct(ctx *gin.Context) {
	productName := ctx.Param("name")

	_, err := handler.client.Database("Products").Collection("Products").InsertOne(context.TODO(), Product{
		ProductName: productName,
	})
	if err != nil {
		panic(err)
	}
}
