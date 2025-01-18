package initialize

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/renxingdawang/rxdw-mall/server/cmd/user/config"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth/authservice"
)

func InitAuth() authservice.Client {
	r, err := consul.NewConsulResolver(fmt.Sprintf("%s:%d",
		config.GlobalConsulConfig.Host,
		config.GlobalConsulConfig.Port))
	if err != nil {
		hlog.Fatalf("new nacos client failed: %s", err.Error())
	}
}
