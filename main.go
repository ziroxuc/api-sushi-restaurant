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
	r.HandleFunc("/pedidosByEstado", actions.GetPedidosPorEstadodEndpoint).Methods("POST")
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



	// Handlers Usuario
	r.HandleFunc("/usuario", actions.CreateUsuarioEndPoint).Methods("POST")
	r.HandleFunc("/usuarios", actions.GetAllUsuariosEndPoint).Methods("GET")
	r.HandleFunc("/usuario/{id}", actions.GetUsuarioByIdEndpoint).Methods("GET")
	r.HandleFunc("/usuario/{id}", actions.UpdateUsuarioEndpoint).Methods("PUT")
	r.HandleFunc("/usuarios/{status}", actions.GetUsuariosPorEstadodEndpoint).Methods("GET")
	r.HandleFunc("/usuario", actions.DeleteUsuarioEndpoint).Methods("DELETE")



	//Login
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/validate", auth.IsValidToken).Methods("GET")



	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:[]string{"Access-Control-Allow-Headers", "Accept", "Content-Type", "Content-Length", "Accept-Encoding"," X-CSRF-Token", "Authorization","X-Requested-With"},
		AllowedMethods:[]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})
	handler := c.Handler(r)

	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatal(err)
	}


}
