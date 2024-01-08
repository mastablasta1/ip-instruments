package main

import (
	"ipapi/ip"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ip", ip.GetIpData) //	Объекты по геопозиции

	//	Запуск сервера на локале
	/*
		log.Println("Запуск веб-сервера на http://127.0.0.1:9990")
		err := http.ListenAndServe(":9990", mux)
		log.Fatal(err)
	*/

	// создадим переменные APP_IP и APP_PORT, а их значения возьмем из одноименных переменных окружения
	APP_IP := os.Getenv("APP_IP")
	APP_PORT := os.Getenv("APP_PORT")

	//	Запустим сервер
	err := http.ListenAndServe(APP_IP+":"+APP_PORT, mux)
	log.Fatal(err)

}
