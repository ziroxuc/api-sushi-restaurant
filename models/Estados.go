package models

import (
	"gopkg.in/mgo.v2/bson"
)

type EstadoObj struct {
	ID             bson.ObjectId `bson:"_id" json:"_id"`
	IdEstado int `bson:"idEstado" json:"idEstado"`
	Nombre string `bson:"nombre" json:"nombre"`
}

type Estados []EstadoObj