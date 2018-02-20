package models

import (
	"gopkg.in/mgo.v2/bson"
	)

type Pedido struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Productos	  `bson:"productos" json:"productos"`
	DatosUsuarios  `bson:"datosUsuario" json:"datosUsuario"`
	Domicilio bool		`bson:"domicilio" json:"domicilio"`
	Fecha_creacion string `bson:"fecha_creacion" json:"fecha_creacion"`
	Estado string `bson:"estado" json:"estado"`
	Total	int	  `bson:"total" json:"total"`
}




