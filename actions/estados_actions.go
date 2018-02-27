package actions

import (
	"net/http"
	mo "../models"
	db "../dbConnection"
	utils "../utils"

)
var cEstados = db.GetCollectionEstados()

func GetAllEstadosEndPoint(w http.ResponseWriter, r *http.Request) {
	var estados []mo.EstadoObj
	err := cEstados.Find(nil).Sort("+idEstado").All(&estados)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion Estados.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,estados)
}

