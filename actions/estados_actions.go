package actions

import (
	"net/http"
	mo "../models"
	db "../dbConnection"

)
var cEstados = db.GetCollectionEstados()

func GetAllEstadosEndPoint(w http.ResponseWriter, r *http.Request) {
	var estados []mo.EstadoObj
	err := cEstados.Find(nil).Sort("+idEstado").All(&estados)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	respondWithJSON(w,http.StatusOK,estados)
}

