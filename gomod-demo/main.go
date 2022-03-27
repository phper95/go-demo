package main

import (
	"gitee.com/phper95/gomod-demo1/pkg/demo1"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Warn("go mod")
	demo1.Demo1()
}
