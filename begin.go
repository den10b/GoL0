package main

import (
	"GoL0/Cache"
	"GoL0/DB"
	"GoL0/HTTP"
	"GoL0/STAN"
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	fmt.Println(uuid.New())
	fmt.Println(uuid.New())
	DB.OpenConn()
	DB.GetAllOrders()
	Cache.CacheInit()
	//Cache.TestCache()
	go HTTP.TestHttp()
	go STAN.TestSub()
	STAN.TestJSONPub()
	//STAN.TestPub()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
