package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/sync/singleflight"
	"sync"
)

var singleFlightGetProduct singleflight.Group

func main() {
	key := "productid"
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(requestID int) {
			defer wg.Done()
			data, err := getProductData(requestID, key)
			if err != nil {
				logs.Error(err)
				return
			}
			logs.Info(data, requestID)
		}(i)
	}
	wg.Wait()
}

func getProductData(requestID int, key string) (interface{}, error) {
	data, _ := getProductDataFromCache(requestID, key)
	if data == nil {
		v, err, _ := singleFlightGetProduct.Do(key, func() (interface{}, error) {
			return getProductDataFromDB(requestID, key)
		})
		return v, err
	}
	return data, nil
}

func getProductDataFromCache(requestID int, key string) (interface{}, error) {
	//只是为了演示，这里可以从缓存服务中获取数据
	fmt.Println("get from cache", requestID)
	return nil, nil
}

func getProductDataFromDB(requestID int, key string) (interface{}, error) {
	//只是为了演示，这里可以从数据存储服务中获取数据
	fmt.Println("get from db requestID", requestID)
	//time.Sleep(time.Second)
	return key, nil
}
