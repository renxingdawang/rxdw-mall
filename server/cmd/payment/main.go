package main

import (
	"log"
	payment "rxdw-mall/server/kitex_gen/payment/paymentservice"
)

func main() {
	svr := payment.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
