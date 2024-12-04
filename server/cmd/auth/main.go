package main

import (
	auth "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth/authservice"
	"log"
)

func main() {
	//IP,Port:=initialize.InitFlag()
	//r,info:=initialize.InitRegistry(Port)

	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
