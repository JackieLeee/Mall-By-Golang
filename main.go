package main

import (
	"flag"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	_ "mall/routers"
)

func main() {
	if err := loadAppConfig(); err != nil {
		logs.Error("Load app conf failed, err: %s", err.Error())
		os.Exit(1)
	}

	initLogger()

	web.Run()
}

func initLogger() {

}

func loadAppConfig() error {
	var conf string

	// 配置文件路径，默认为conf/app.conf
	flag.StringVar(&conf, "c", "conf/app.conf", "config file")
	// 解析参数
	flag.Parse()

	if err := web.LoadAppConfig("ini", conf); err != nil {
		return err
	}
	return nil
}
