package initialize

import (
	"flag"

	"gin_start_demo/global"
	"gin_start_demo/model"
)

var (
	port     int
	coreHost string
	corePort int
)

func InitConfig() {
	flag.IntVar(&port, "port", 21103, "set port")
	flag.Parse()
	conf := model.Config{
		Port: port,
	}
	global.Config = &conf
}
