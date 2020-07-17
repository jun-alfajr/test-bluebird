package product

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/jun-alfajr/test-bluebird/db"
)

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
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("product")
	err = c.Insert(product)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(product)
}

func listProducts(req *restful.Request, resp *restful.Response) {
	allProducts := make([]Product, 0)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("product")
	err := c.Find(bson.M{}).All(&allProducts)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(allProducts)
}

func getProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("productId")
	product := Product{}
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("product")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(productId)}).One(&product)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(product)
}

func updateProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("productId")
	product := Product{}
	err := req.ReadEntity(&product)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}

	product.ID = bson.ObjectIdHex(productId)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("product")
	err = c.Update(bson.M{"_id": product.ID}, product)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(product)
}

func deleteProduct(req *restful.Request, resp *restful.Response) {
	productId := req.PathParameter("productId")

	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("product")
	err := c.RemoveId(productId)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteHeader(200)
}
