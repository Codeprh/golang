package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func startHttp() {
	http.HandleFunc("/healthz", healthz)
	//Use the default DefaultServeMux.
	//ListenAndService
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//应用健康监听端口
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func main() {
	fmt.Println("start")
	startHttp()
}
