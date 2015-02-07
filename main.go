package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	
	"github.com/gorilla/mux"
	"etcd-dashboard/dashboard"
	"etcd-dashboard/config"
)

func main() {
	
	var config = config.New()
	
	if err := config.LoadFlags(os.Args[1:]); err != nil {
		fmt.Println(err.Error() + "\n")
		os.Exit(1)	
	}
	
	r := mux.NewRouter()
	r.PathPrefix("/mod").Handler(http.StripPrefix("/mod", dashboard.HTTPHandler()))
	
	http.Handle("/", r)
	var addr string = config.Host + ":" + config.Port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}								
}