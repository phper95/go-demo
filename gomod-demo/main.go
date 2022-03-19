package main

import "go.uber.org/zap"
import "gitee.com/phper95/go-demo/gomod-demo1/pkg/demo1"

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Error("test error")
	demo1.Demo1()
}
