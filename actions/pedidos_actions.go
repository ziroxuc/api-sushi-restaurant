package actions

import (
	"net/http"
	"encoding/json"
	mo "../models"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gorilla/mux"
	"log"
	"strconv"
	db "../dbConnection"

)


var cPedidos = db.GetCollectionPedidos()

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllPedidosEndPoint(w http.ResponseWriter, r *http.Request) {
	var pedidos []mo.Pedido
	err := cPedidos.Find(nil).Sort("+id_sushi").All(&pedidos)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	respondWithJSON(w,http.StatusOK,pedidos)
}

func CreatePedidoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pedido mo.Pedido

	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error al leer el body.")
		return
	}
	pedido.ID = bson.NewObjectId()
	t := time.Now()
	var timeMod = t.Format("02-01-2006 15:04:05")
	pedido.Fecha_creacion = timeMod

	if err := cPedidos.Insert(pedido); err != nil {
		respondWithError(w, http.StatusBadRequest, "No se pudo insertar el resgistro.")
		return
	}
	respondWithJSON(w,http.StatusCreated,pedido)
}

func FindPedidoEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoId := params["id"]

	var resultado mo.Pedido
	idConv,_ := strconv.Atoi(pedidoId);
	err := cPedidos.Find(bson.M{"id":idConv}).One(&resultado)
	if (err != nil) {
		respondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	respondWithJSON(w,http.StatusOK,resultado)
}

func FindPedidoIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoId := params["id"]
	 var resultado mo.Pedido

	if !bson.IsObjectIdHex(pedidoId){
		respondWithError(w, http.StatusInternalServerError, "No es un ObjectIDHex.")
		return
	}
	err := cPedidos.FindId(bson.ObjectIdHex(pedidoId)).One(&resultado)
	if (err != nil) {
		respondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	respondWithJSON(w,http.StatusOK,resultado)
}

func UpdatePedidoEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pedidoId := params["id"]

	if !bson.IsObjectIdHex(pedidoId){
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var pedido_data mo.Pedido

	err := decoder.Decode(&pedido_data)
	if (err != nil) {
		respondWithError(w, http.StatusInternalServerError, "Error al enviar json de productos.")
		return
	}
	defer r.Body.Close()

	idDeserialize := bson.ObjectIdHex(pedidoId);

	document := bson.M{"_id":idDeserialize}
	change := bson.M{"$set":pedido_data}

	error := cPedidos.Update(document,change)

	if(error != nil){
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedido_data)
}

func GetPedidosPorEstadodEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["status"]
	var pedidos []mo.Pedido

	var idEstado,erro = strconv.Atoi(status)
	if erro!= nil{
		respondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cPedidos.Find(bson.M{"estado":idEstado}).All(&pedidos)
	if (err != nil) {
		respondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	respondWithJSON(w,http.StatusOK,pedidos)
}

