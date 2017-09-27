package main

import (
	"fmt"
	"meuse/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/listproduct", controller.GetListProduct).Methods("GET")
	r.HandleFunc("/openbox", controller.OpenBox).Methods("GET")
	r.HandleFunc("/allocatebox", controller.AllocateBox).Methods("POST")

	fmt.Println("Server em pé")
	http.ListenAndServe(":8085", r)
}
