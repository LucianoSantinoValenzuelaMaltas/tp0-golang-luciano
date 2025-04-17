package main

import (
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()
	//http.ListenAndServe(": 8080", mux)

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	//panic("no implementado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

	//utils.RecibirMensaje()
	//utils.RecibirPaquetes()
}
