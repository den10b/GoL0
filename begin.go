package main

import (
	"GoL0/Cache"
	"GoL0/DB"
	"GoL0/HTTP"
	"GoL0/STAN"
	"bufio"
	"log"
	"os"
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

	log.Printf("Нажмите enter для отправки файла в канал")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	STAN.TestJSONPub("testJSON.json")

	log.Printf("Нажмите enter для завершения работы сервиса")
	input = bufio.NewScanner(os.Stdin)
	input.Scan()

	STAN.QuitSub()
	DB.CloseConn()

}
