package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	res1 := ""
	res2 := ""
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		res1 = searchThread1(&wg)

	}()

	go func() {
		res2 = searchThread2(&wg)
	}()

	wg.Wait()
	fmt.Println("all search done", res1, res2)

}

func searchThread1(wg *sync.WaitGroup) string {
	defer func() {
		wg.Done()
	}()
	//do search t1
	time.Sleep(time.Second)
	fmt.Println("t1 search")
	return "t1"

}

func searchThread2(wg *sync.WaitGroup) string {
	defer func() {
		wg.Done()
	}()
	//do search t2
	time.Sleep(time.Second)
	fmt.Println("t2 search")
	return "t2"

}
