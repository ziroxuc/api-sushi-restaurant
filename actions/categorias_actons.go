package actions

import (
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	utils "../utils"
	mo "../models"
	"github.com/gorilla/mux"
	db "../dbConnection"
	"fmt"
)

var cCategoria = db.GetCollectionCategorias()

func CreateCategoriaEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var categoria mo.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer el body.")
		return
	}
	categoria.ID = bson.NewObjectId();

	if err := cCategoria.Insert(categoria); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No se pudo insertar el resgistro.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated,categoria)
}

func GetAllCategoriasEndPoint(w http.ResponseWriter, r *http.Request) {
	var categorias mo.Categorias
	err := cCategoria.Find(nil).Sort("+nombre").All(&categorias)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,categorias)
}

func GetCategoriaByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCate := params["id"]
	var categoria mo.Categoria

	if !bson.IsObjectIdHex(idCate){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	err := cCategoria.FindId(bson.ObjectIdHex(idCate)).One(&categoria)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,categoria)
}

func UpdateCategoriaEndpoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	categoriaID := params["id"]

	if !bson.IsObjectIdHex(categoriaID){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var categoria_data mo.Categoria

	err := decoder.Decode(&categoria_data)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error en el envio de parametros entrada de Categoria.")
		return
	}
	defer r.Body.Close()
	idDeserialize := bson.ObjectIdHex(categoriaID);

	document := bson.M{"_id":idDeserialize}
	change := bson.M{"$set":categoria_data}

	error := cCategoria.Update(document,change)

	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al actualizar categoria.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,categoria_data)
}

func DeleteCategoriaEndpoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var categoria mo.Categoria
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	error := cCategoria.Remove(categoria)
	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar producto.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,map[string]string{"result": "Categoria eliminada"})
}