package actions

import (
	"net/http"
	mo "../models"
	db "../dbConnection"
	utils "../utils"
	auth "../authentication"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)
var cEstados = db.GetCollectionEstados()

func CreateEstadoEndPoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	defer r.Body.Close()
	var estado mo.EstadoObj
	if err := json.NewDecoder(r.Body).Decode(&estado); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer el body.")
		return
	}
	estado.ID = bson.NewObjectId();

	if err := cEstados.Insert(estado); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No se pudo insertar el resgistro.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated, estado)
}

func GetAllEstadosEndPoint(w http.ResponseWriter, r *http.Request) {
	var estados []mo.EstadoObj
	err := cEstados.Find(nil).Sort("+idEstado").All(&estados)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion Estados.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,estados)
}


