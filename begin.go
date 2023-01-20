package main

import (
	"GoL0/Cache"
	"GoL0/DB"
	"GoL0/HTTP"
	"GoL0/STAN"
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	err := DB.OpenConn()
	if err != nil {
		log.Printf("Error initing DB connection: %s\n", err)
		return
	}
	err = Cache.CacheInit()
	if err != nil {
		log.Printf("Error initing cache: %s\n", err)
		return
	}
	//Cache.TestCache()
	go HTTP.InitHttp()
	go STAN.InitSub()
	time.Sleep(3 * time.Second)
	STAN.TestJSONPub("testJSON.json")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	STAN.QuitSub()
	DB.CloseConn()

}
