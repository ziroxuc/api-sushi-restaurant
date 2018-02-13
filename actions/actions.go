package actions

import (
	"net/http"
	"encoding/json"
	mo "../models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gorilla/mux"
	"log"
	"strconv"
)

func getSession() *mgo.Session {
	session, error := mgo.Dial("mongodb://localhost")
	if (error != nil) {
		panic(error)
	}
	return session
}

var collection = getSession().DB("sushi-restaurant-db").C("pedidos")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor funciona")
}

func AllPedidosEndPoint(w http.ResponseWriter, r *http.Request) {

	var pedidos []mo.Pedido
	err := collection.Find(nil).Sort("+id_sushi").All(&pedidos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedidos)
}

func CreatePedidoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pedido mo.Pedido
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {

		fmt.Println(err)
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pedido.ID = bson.NewObjectId()
	t := time.Now()
	var timeMod = t.Format("02-01-2006 15:04:05")
	pedido.Fecha_creacion = timeMod

	if err := collection.Insert(pedido); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pedido)
}

func FindPedidoEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoId := params["id"]

	var resultado mo.Pedido

	idConv,_ := strconv.Atoi(pedidoId);
	err := collection.Find(bson.M{"id_sushi":idConv}).One(&resultado)
	if (err != nil) {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
