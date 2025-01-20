package initialize

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/tools"
)

// InitFlag to init flag
// InitFlag 方法的作用是 初始化命令行参数，并返回解析后的 IP 地址和端口号。以下是该方法的详细解释：
func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)
	// Parsing flags and if Port is 0 , then will automatically get an empty Port.
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	klog.Info("ip: ", *IP)
	klog.Info("port: ", *Port)
	return *IP, *Port
}
