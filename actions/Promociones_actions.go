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
	auth "../authentication"
)

var cPromocion = db.GetCollectionPromociones()

func CreatePromocionEndPoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	defer r.Body.Close()
	var promocion mo.Promocion
	if err := json.NewDecoder(r.Body).Decode(&promocion); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al decodificar parametros de entrada.")
		return
	}
	promocion.ID = bson.NewObjectId()
	promocion.Create_date = time.Now()

	if err := cPromocion.Insert(promocion); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al crear promocion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated, promocion)
}

func GetAllPromocionesEndPoint(w http.ResponseWriter, r *http.Request) {
	var promociones []mo.Promocion
	err := cPromocion.Find(nil).Sort("+id_sushi").All(&promociones)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,promociones)
}

func UpdatePromocionEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	params := mux.Vars(r)
	PromocionID := params["id"]

	if !bson.IsObjectIdHex(PromocionID){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var Promocion_data mo.Promocion

	err := decoder.Decode(&Promocion_data)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error en el envio de parametros entrada de Promocion.")
		return
	}
	defer r.Body.Close()

	idDeserialize := bson.ObjectIdHex(PromocionID);
	Promocion_data.Mod_date = time.Now()
	Promocion_data.Create_date = Promocion_data.Create_date

	document := bson.M{"_id":idDeserialize}
	change := bson.M{"$set":Promocion_data}

	error := cPromocion.Update(document,change)

	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al actualizar Promocion.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,Promocion_data)
}

func GetPromocionesPorEstadodEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["status"]
	var Promociones []mo.Promocion

	var idEstado,erro = strconv.Atoi(status)
	if erro!= nil{
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cPromocion.Find(bson.M{"estado":idEstado}).All(&Promociones)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,Promociones)
}

func DeletePromocionEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	defer r.Body.Close()
	var Promocion mo.Promocion
	if err := json.NewDecoder(r.Body).Decode(&Promocion); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	error := cPromocion.Remove(Promocion)
	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar Promocion.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,map[string]string{"result": "Promocion eliminada"})
}

func GetPromocionByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProd := params["id"]
	var Promocion mo.Promocion

	if !bson.IsObjectIdHex(idProd){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	err := cPromocion.FindId(bson.ObjectIdHex(idProd)).One(&Promocion)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,Promocion)
}