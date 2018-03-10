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

var cUsuario = db.GetCollectionUsuario()

func CreateUsuarioEndPoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	defer r.Body.Close()
	var usuario mo.User
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al decodificar parametros de entrada.")
		return
	}
	usuario.ID = bson.NewObjectId()
	usuario.Create_date = time.Now()
	//hashea password
	passHasheada := utils.HashearPassword(usuario.Password)

	usuario.PasswordHash = passHasheada
	usuario.Password = ""
	if err := cUsuario.Insert(usuario); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al crear Usuario.")
		return
	}
	utils.RespondWithJSON(w,http.StatusCreated,usuario)
}

func GetAllUsuariosEndPoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	var usuarios []mo.User
	err := cUsuario.Find(nil).Sort("+nombre").All(&usuarios)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "No se encontraron registros en la coleccion.")
		return
	}
	utils.RespondWithJSON(w,http.StatusOK,usuarios)
}

func UpdateUsuarioEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	params := mux.Vars(r)
	UsuarioID := params["id"]

	if !bson.IsObjectIdHex(UsuarioID){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var Usuario_data mo.User

	err := decoder.Decode(&Usuario_data)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error en el envio de parametros entrada de Usuario.")
		return
	}
	defer r.Body.Close()

	idDeserialize := bson.ObjectIdHex(UsuarioID);
	Usuario_data.Mod_date = time.Now()
	Usuario_data.Create_date = Usuario_data.Create_date

	document := bson.M{"_id":idDeserialize}
	change := bson.M{"$set":Usuario_data}

	error := cUsuario.Update(document,change)

	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al actualizar Usuario.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,Usuario_data)
}

func GetUsuariosPorEstadodEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	params := mux.Vars(r)
	status := params["status"]
	var Usuarios []mo.User

	var idEstado,erro = strconv.Atoi(status)
	if erro!= nil{
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al convertir parametro int.")
	}
	err := cUsuario.Find(bson.M{"estado":idEstado}).All(&Usuarios)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,Usuarios)
}

func DeleteUsuarioEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	defer r.Body.Close()
	var Usuario mo.User
	if err := json.NewDecoder(r.Body).Decode(&Usuario); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	error := cUsuario.Remove(Usuario)
	if(error != nil){
		fmt.Println(error)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar Usuario.")
		return
	}
	utils.RespondWithJSON(w ,http.StatusOK ,map[string]string{"result": "success"})
}

func GetUsuarioByIdEndpoint(w http.ResponseWriter, r *http.Request) {

	isAuth := auth.ValidateToken(r)
	if(isAuth != ""){
		utils.RespondWithError(w, http.StatusUnauthorized, isAuth)
		return
	}

	params := mux.Vars(r)
	idProd := params["id"]
	var Usuario mo.User

	if !bson.IsObjectIdHex(idProd){
		utils.RespondWithError(w, http.StatusBadRequest, "El id no es un ObjectIdHex valido.")
		return
	}
	err := cUsuario.FindId(bson.ObjectIdHex(idProd)).One(&Usuario)
	if (err != nil) {
		utils.RespondWithError(w, http.StatusNotFound, "No se encontró el registro.")
	}
	utils.RespondWithJSON(w,http.StatusOK,Usuario)
}