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
	utils "../utils"
)

var cPedidos = db.GetCollectionPedidos()

func AllPedidosEndPoint(w http.ResponseWriter, r *http.Request) {
	var pedidos []mo.Pedido
	err := cPedidos.Find(nil).Sort("+id_sushi").All(&pedidos)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,pedidos)
}

func CreatePedidoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var pedido mo.Pedido

	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer el body.")
		return
	}
	pedido.ID = bson.NewObjectId()
	t := time.Now()
	var timeMod = t.Format("02-01-2006 15:04:05")
	pedido.Fecha_creacion = timeMod

	if err := cPedidos.Insert(pedido); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No se pudo insertar el resgistro.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated,pedido)
}

func FindPedidoEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoId := params["id"]

	var resultado mo.Pedido
	idConv,_ := strconv.Atoi(pedidoId);
	err := cPedidos.Find(bson.M{"id":idConv}).One(&resultado)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,resultado)
}

func FindPedidoIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pedidoId := params["id"]
	 var resultado mo.Pedido

	if !bson.IsObjectIdHex(pedidoId){
		utils.RespondWithError(w, http.StatusInternalServerError, "No es un ObjectIDHex.")
		return
	}
	err := cPedidos.FindId(bson.ObjectIdHex(pedidoId)).One(&resultado)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,resultado)
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
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al enviar json de pedido.")
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
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cPedidos.Find(bson.M{"estado":idEstado}).All(&pedidos)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,pedidos)
}

func GetCantRegistrosByEstadosEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var estados []mo.EstadoObj
	var cantPedidos []int
	suma := 0

	if err := json.NewDecoder(r.Body).Decode(&estados); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer parametros de entrada.")
		return
	}
	for _, valor := range estados {
		cantidadReg,err := cPedidos.Find(bson.M{"estado": valor.IdEstado}).Count()
		if(err!= nil){
			utils.RespondWithError(w, http.StatusInternalServerError, "Error al contar regstros del id: "+valor.Nombre)
			return
		}
		cantPedidos = append(cantPedidos, cantidadReg)
		suma += cantidadReg

	}
	cantPedidos = append(cantPedidos, suma)

	utils.RespondWithJSON(w,http.StatusOK,cantPedidos)
}




