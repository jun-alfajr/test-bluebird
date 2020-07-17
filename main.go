package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jun-alfajr/test-bluebird/product"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	wsContainer := restful.NewContainer()
	wsContainer.Add(product.ProductController{}.AddRouters())

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	host := "127.0.0.1:8080"
	log.Printf("listening on: %s", host)
	server := &http.Server{Addr: host, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
