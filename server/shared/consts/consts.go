package consts

const (
	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName    = "port"
	PortFlagUsage   = "port"
	FreePortAddress = "localhost:0"
	CorsAddress     = "http://localhost:3000"
	TCP             = "tcp"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"
)
