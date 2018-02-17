package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"./actions"

	"github.com/rs/cors"
)
func main(){
	r := mux.NewRouter()

	r.HandleFunc("/pedido/{id}", actions.FindPedidoEndpoint).Methods("GET")
	r.HandleFunc("/pedidos", actions.AllPedidosEndPoint).Methods("GET")
	r.HandleFunc("/pedido", actions.CreatePedidoEndPoint).Methods("POST")
	//r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	//r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")


	// handlers productos
	r.HandleFunc("/producto", actions.CreateProductoEndPoint).Methods("POST")


	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200","http://localhost:4300"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatal(err)
	}


}
