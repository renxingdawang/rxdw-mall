package main

import (
	payment "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment/paymentservice"
	"log"
)

func main() {
	svr := payment.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
