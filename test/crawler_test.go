package test

import (
	"fmt"
	"github.com/codecodify/news/lib/impl"
	"testing"
)

func TestFetchZhihu(t *testing.T) {
	client := impl.Zhihu{}
	response, err := client.Get(0)
	fmt.Println("err", err)
	fmt.Println(response)
}

func TestFetchWangyi(t *testing.T) {
	client := impl.Wangyi{}
	response, err := client.Get(0)
	fmt.Println("err", err)
	fmt.Printf("%+v", response.Data.Date)
}
