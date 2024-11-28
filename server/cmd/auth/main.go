package main

import (
	"log"
	auth "rxdw-mall/server/kitex_gen/auth/authservice"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
