package main

import (
	"log"
	order "rxdw-mall/server/kitex_gen/order/orderservice"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
