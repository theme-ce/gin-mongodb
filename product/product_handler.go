package product

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type productHandler struct {
	client *mongo.Client
}

type ProductHandler interface {
	InsertProduct(ctx *gin.Context)
}
