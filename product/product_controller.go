package product

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo/bson"
)

// in memory storage
var products map[string]Product = make(map[string]Product)

type ProductController struct {
}

func (controller ProductController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/product").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/").To(createProduct))
	ws.Route(ws.GET("/").To(listProducts))
	ws.Route(ws.GET("/{productId}").To(getProduct))
	ws.Route(ws.PUT("/{productId}").To(updateProduct))
	ws.Route(ws.DELETE("/{productId}").To(deleteProduct))

	return ws
}

func createProduct(req *restful.Request, resp *restful.Response) {
	product := Product{}
	err := req.ReadEntity(&product)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	product.ID = bson.NewObjectId()
	products[product.ID.Hex()] = product
	resp.WriteEntity(product)
}

func listProducts(req *restful.Request, resp *restful.Response) {
	allProducts := make([]Product, 0)
	for _, product := range products {
		allProducts = append(allProducts, product)
	}
	resp.WriteEntity(allProducts)
}

func getProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("producId")
	if _, ok := products[productId]; !ok {
		resp.WriteHeaderAndEntity(404, "not found")
		return
	}

	resp.WriteEntity(products[productId])
}

func updateProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("productId")
	product := Product{}
	err := req.ReadEntity(&product)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}
	if _, ok := products[productId]; !ok {
		resp.WriteHeaderAndEntity(404, "not found")
		return
	}

	product.ID = bson.ObjectIdHex(productId)
	products[product.ID.Hex()] = product
	resp.WriteEntity(product)
}

func deleteProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("productId")

	if _, ok := products[productId]; !ok {
		resp.WriteHeaderAndEntity(404, "not found")
		return
	}

	delete(products, productId)
	resp.WriteHeader(200)
}
