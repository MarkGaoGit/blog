package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href='http://localhost:8009/list'>seelog日志</a>")
}

func queryList(w http.ResponseWriter, r *http.Request) {

}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		defaultPage(w, r)
		break
	case "/list":
		queryList(w, r)
		break
	}

}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":8009", mux)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}

}
