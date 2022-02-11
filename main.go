package main

import (
	"fmt"
	"net"
	"sort"
)

//设置扫描ip
var str = "127.0.0.1"

//c1：扫描的端口号，c2：能连接的端口
func work(c1, c2 chan int, str string) {
	for ch := range c1 {
		address := fmt.Sprintf("%s:%d", str, ch)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			c2 <- 0
			continue
		}
		err = conn.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println(ch, "端口打开了")
		c2 <- ch
	}
}

func main() {
	//通道缓存1000
	c := make(chan int, 1000)
	res := make(chan int)
	var op []int
	for i := 0; i < cap(c); i++ {
		go work(c, res, str)
	}
	go func() {
		for i := 21; i < 32769; i++ {
			c <- i
		}
	}()
	for i := 21; i < 32769; i++ {
		p := <-res
		if p != 0 {
			op = append(op, p)
		}
	}

	sort.Ints(op)
	fmt.Printf("%v所有打开的端口包括%v", str, op)
	close(c)
	close(res)
}
