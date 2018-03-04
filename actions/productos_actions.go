package actions

import (
	"net/http"
	"encoding/json"
	"time"
	db "../dbConnection"
	mo "../models"
	utils "../utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strconv"
)

var cProducto = db.GetCollectionProductos()

func CreateProductoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var producto mo.Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al decodificar parametros de entrada.")
		return
	}
	producto.ID = bson.NewObjectId()
	producto.Create_date = time.Now()

	if err := cProducto.Insert(producto); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al crear producto.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated,producto)
}

func GetAllProductosEndPoint(w http.ResponseWriter, r *http.Request) {
	var prductos []mo.Producto
	err := cProducto.Find(nil).Sort("+id_sushi").All(&prductos)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,prductos)
}

func UpdateProductoEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	productoID := params["id"]

	if !bson.IsObjectIdHex(productoID){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var producto_data mo.Producto

	err := decoder.Decode(&producto_data)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error en el envio de parametros entrada de Producto.")
		return
	}
	defer r.Body.Close()

	idDeserialize := bson.ObjectIdHex(productoID);
	producto_data.Mod_date = time.Now()
	producto_data.Create_date = producto_data.Create_date

	document := bson.M{"_id":idDeserialize}
	change := bson.M{"$set":producto_data}

	error := cProducto.Update(document,change)

	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al actualizar producto.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,producto_data)
}

func GetProductosPorEstadodEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["status"]
	var productos []mo.Producto

	var idEstado,erro = strconv.Atoi(status)
	if erro!= nil{
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cProducto.Find(bson.M{"estado":idEstado}).All(&productos)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,productos)
}

func GetProductosPorCategoriaEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]
	var productos []mo.Producto

	var idCategory,erro = strconv.Atoi(category)
	if erro!= nil{
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cProducto.Find(bson.M{"id_categoria":idCategory}).All(&productos)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontrarón regstros.")
	}
	utils.RespondWithJSON(w,http.StatusOK,productos)
}

func DeleteProductoEndpoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var producto mo.Producto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	error := cProducto.Remove(producto)
	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar producto.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,map[string]string{"result": "success"})
}

func GetProductoByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProd := params["id"]
	var producto mo.Producto

	if !bson.IsObjectIdHex(idProd){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	err := cProducto.FindId(bson.ObjectIdHex(idProd)).One(&producto)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,producto)
}