package main

import (
	"github.com/zouhuigang/package/zlog"
)

func main() {
	zlog.Init("./", "INFO")
	zlog.Infof("test sucess log")
}
