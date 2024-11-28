// Code generated by Kitex v0.11.3. DO NOT EDIT.
package authservice

import (
	server "github.com/cloudwego/kitex/server"
	"github.com/rxdw-mall/server/shared/kitex_gen/auth"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler auth.AuthService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler auth.AuthService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
