package main

import (
	checkout "github.com/rxdw-mall/server/shared/kitex_gen/checkout/checkoutservice"
	"log"
)

func main() {
	svr := checkout.NewServer(new(CheckoutServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
