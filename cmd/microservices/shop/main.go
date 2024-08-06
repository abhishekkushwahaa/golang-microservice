package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abhishekkushwahaa/golang-microservice/pkg/common/cmd"
	"github.com/abhishekkushwahaa/golang-microservice/pkg/shop"
	shop_app "github.com/abhishekkushwahaa/golang-microservice/pkg/shop/application"
	shop_infra_product "github.com/abhishekkushwahaa/golang-microservice/pkg/shop/infrastructure/products"
	shop_interfaces_private_http "github.com/abhishekkushwahaa/golang-microservice/pkg/shop/interfaces/private/http"
	shop_interfaces_public_http "github.com/abhishekkushwahaa/golang-microservice/pkg/shop/interfaces/public/http"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the shop microservices!")

	ctx := cmd.Context()

	r := createShopMicrodervice()

	server := &http.Server{Addr: os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()

	log.Println("Closing shop microservice!")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createShopMicrodervice() *chi.Mux {
	shopProductRepo := shop_infra_product.NewMemoryRepository()
	shopProductsService := shop_app.NewProductsService(shopProductRepo, shopProductRepo)

	if err := shop.LoadShopFixtures(shopProductsService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)

	return r
}
