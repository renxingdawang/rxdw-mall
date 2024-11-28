package main

import (
	"log"
	checkout "rxdw-mall/server/kitex_gen/checkout/checkoutservice"
)

func main() {
	svr := checkout.NewServer(new(CheckoutServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
