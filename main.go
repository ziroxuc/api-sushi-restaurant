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

	//r.HandleFunc("/pedido/{id}", actions.FindPedidoEndpoint).Methods("GET")
	r.HandleFunc("/pedido/{id}", actions.FindPedidoIdEndpoint).Methods("GET")
	r.HandleFunc("/pedidos", actions.AllPedidosEndPoint).Methods("GET")
	r.HandleFunc("/pedido", actions.CreatePedidoEndPoint).Methods("POST")
	r.HandleFunc("/pedido/{id}", actions.UpdatePedidoEndpoint).Methods("PUT")
	r.HandleFunc("/pedidoByEstado/{status}", actions.GetPedidosPorEstadodEndpoint).Methods("GET")
	r.HandleFunc("/pedidosCount", actions.GetCantRegistrosByEstadosEndpoint).Methods("POST")


	//r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	//r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

	// Handlers estados
	r.HandleFunc("/estados", actions.GetAllEstadosEndPoint).Methods("GET")



	// Handlers productos
	r.HandleFunc("/producto", actions.CreateProductoEndPoint).Methods("POST")
	r.HandleFunc("/productos", actions.GetAllProductosEndPoint).Methods("GET")
	r.HandleFunc("/producto/{id}", actions.GetProductoByIdEndpoint).Methods("GET")
	r.HandleFunc("/producto/{id}", actions.UpdateProductoEndpoint).Methods("PUT")
	r.HandleFunc("/productos/{status}", actions.GetProductosPorEstadodEndpoint).Methods("GET")
	r.HandleFunc("/productosCat/{category}", actions.GetProductosPorCategoriaEndpoint).Methods("GET")
	r.HandleFunc("/producto", actions.DeleteProductoEndpoint).Methods("DELETE")




	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200","http://localhost:4300"},
		AllowCredentials: true,
		AllowedHeaders:[]string{"X-Requested-With","content-type"},
		AllowedMethods:[]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})
	handler := c.Handler(r)

	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatal(err)
	}


}
