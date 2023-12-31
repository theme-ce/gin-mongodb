package product

import "github.com/theme-ce/gin-mongodb/app"

type productStorer interface {
	NewProduct(*Product)
	GetAllProduct() []Product
}

type ProductHandler struct {
	storer productStorer
}

func NewProductHandler(store productStorer) *ProductHandler {
	return &ProductHandler{storer: store}
}

func (p *ProductHandler) InsertProduct(ctx app.Context) {
	product := Product{
		ProductName: ctx.Params("name"),
	}
	p.storer.NewProduct(&product)
}

func (p *ProductHandler) GetAllProduct(ctx app.Context) {
	products := p.storer.GetAllProduct()
	ctx.OK(products)
}
