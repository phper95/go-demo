package main

import (
	"fmt"
	"time"
)

var t1 chan bool
var t2 chan bool

func main() {
	res1 := ""
	res2 := ""
	t1 = make(chan bool, 0)
	t2 = make(chan bool, 0)
	go func() {
		res1 = searchThread1()
	}()

	go func() {
		res2 = searchThread2()
	}()

	//一旦接收到消息就会执行，否则阻塞等待
	<-t1
	<-t2
	fmt.Println("all search done", res1, res2)
}

func searchThread1() string {
	defer func() {
		t1 <- true
	}()
	//do search t1
	time.Sleep(time.Second)
	fmt.Println("t1 search")
	//执行完成后发送消息到管道t1
	return "t1"

}

func searchThread2() string {
	defer func() {
		t2 <- true
	}()
	//do search t2
	time.Sleep(time.Second)
	fmt.Println("t2 search")
	//执行完成后发送消息到管道t2
	return "t2"

}
