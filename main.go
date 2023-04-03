package main

import (
	"github.com/haolie/goUtil/logUtil"
	"lyh/goBase/config"
	"lyh/goBase/sys"
)

func main() {
	config.LoadConfig()
	logUtil.InitLog()

	ctx, success := sys.Start()
	if success == false {
		<-ctx.Done()
	}
}
