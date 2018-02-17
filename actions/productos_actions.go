package actions

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"time"
	db "../dbConnection"
	mo "../models"
)

var cProducto = db.GetCollectionProductos()

func CreateProductoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var producto mo.Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		fmt.Println(err)
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t := time.Now()
	var timeMod = t.Format("02-01-2006 15:04:05")
	producto.Create_date = timeMod

	if err := cProducto.Insert(producto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(producto)
}