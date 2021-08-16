package main

import (
	"github.com/cwowhappy/study-golang/studypprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add(""))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
