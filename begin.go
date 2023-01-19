package main

import (
	"GoL0/Cache"
	"GoL0/DB"
	"GoL0/HTTP"
)

func main() {
	DB.OpenConn()
	Cache.CacheInit()
	//Cache.TestCache()
	HTTP.TestHttp()
	//STAN.TestSub()
	//STAN.TestPub()

}
