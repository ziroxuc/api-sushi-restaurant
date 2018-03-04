package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Promocion struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Id_promocion	int			  `bson:"id_promocion" json:"id_promocion"`
	Nombre      	string        `bson:"nombre" json:"nombre"`
	Precio  		int        	  `bson:"precio" json:"precio"`
	Precio_antiguo  int        	  `bson:"precio_antiguo" json:"precio_antiguo"`
	Piezas      	int        	  `bson:"piezas" json:"piezas"`
	Descripcion 	string        `bson:"descripcion" json:"descripcion"`
	Img 			string        `bson:"img" json:"img"`
	Create_date 	time.Time      `bson:"create_date" json:"create_date"`
	Mod_date 		time.Time      `bson:"mod_date" json:"mod_date"`
	Estado 			int           `bson:"estado" json:"estado"`
	Cantidad 		int			  `bson:"cantidad,omitempty" json:"cantidad,omitempty"`
	Destacado		bool		  `bson:"destacado,omitempty" json:"destacado,omitempty"`
	Stock			bool		  `bson:"stock,omitempty" json:"stock,omitempty"`
	Productos	  `bson:"productos" json:"productos"`
}

type Promociones []Promocion
