package main

import (
	product "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/product/productcatalogservice"
	"log"
)

func main() {
	svr := product.NewServer(new(ProductCatalogServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
