package initialize

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/tools"
)

func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)

	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	klog.Info("ip:", *IP)
	klog.Info("port:", *Port)
	return *IP, *Port
}
