package models

import "gopkg.in/mgo.v2/bson"

type Categoria struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Id_categoria  	int        	  	`bson:"id_categoria" json:"id_categoria"`
	Nombre      	string        	  	`bson:"nombre" json:"nombre"`
}
type Categorias []Categoria

