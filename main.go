package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"./actions"
	auth "./authentication"

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

	// Handlers estados
	r.HandleFunc("/estado", actions.CreateEstadoEndPoint).Methods("POST")
	r.HandleFunc("/estados", actions.GetAllEstadosEndPoint).Methods("GET")


	// Handlers productos
	r.HandleFunc("/producto", actions.CreateProductoEndPoint).Methods("POST")
	r.HandleFunc("/productos", actions.GetAllProductosEndPoint).Methods("GET")
	r.HandleFunc("/producto/{id}", actions.GetProductoByIdEndpoint).Methods("GET")
	r.HandleFunc("/producto/{id}", actions.UpdateProductoEndpoint).Methods("PUT")
	r.HandleFunc("/productos/{status}", actions.GetProductosPorEstadodEndpoint).Methods("GET")
	r.HandleFunc("/productosCat/{category}", actions.GetProductosPorCategoriaEndpoint).Methods("GET")
	r.HandleFunc("/producto", actions.DeleteProductoEndpoint).Methods("DELETE")


	// Handlers Promociones
	r.HandleFunc("/promocion", actions.CreatePromocionEndPoint).Methods("POST")
	r.HandleFunc("/promociones", actions.GetAllPromocionesEndPoint).Methods("GET")
	r.HandleFunc("/promocion/{id}", actions.GetPromocionByIdEndpoint).Methods("GET")
	r.HandleFunc("/promocion/{id}", actions.UpdatePromocionEndpoint).Methods("PUT")
	r.HandleFunc("/promociones/{status}", actions.GetPromocionesPorEstadodEndpoint).Methods("GET")
	r.HandleFunc("/promocion", actions.DeletePromocionEndpoint).Methods("DELETE")


	//Categorias
	r.HandleFunc("/categoria", actions.CreateCategoriaEndPoint).Methods("POST")
	r.HandleFunc("/categorias", actions.GetAllCategoriasEndPoint).Methods("GET")
	r.HandleFunc("/categoria/{id}", actions.GetCategoriaByIdEndpoint).Methods("GET")
	r.HandleFunc("/categoria/{id}", actions.UpdateCategoriaEndpoint).Methods("PUT")
	r.HandleFunc("/categoria", actions.DeleteCategoriaEndpoint).Methods("DELETE")


	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/validate", auth.ValidateToken).Methods("GET")



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
