package main

import (
	"errors"
	"fmt"
	pkgErr "github.com/pkg/errors"
)

var originErr = errors.New("origin error")

func main() {
	err := testErr3()
	//fmt.Printf("%+v", err)
	causeErr := pkgErr.Cause(err)
	fmt.Printf("%+v", causeErr)
}

func testErr1() error {
	return pkgErr.Wrapf(originErr, "add  testErr1")
}

func testErr2() error {
	return testErr1()
}

func testErr3() error {
	return testErr2()
}
