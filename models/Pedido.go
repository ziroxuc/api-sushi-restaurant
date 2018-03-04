package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Pedido struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Productos	  `bson:"productos" json:"productos"`
	DatosUsuario  `bson:"datosUsuario" json:"datosUsuario"`
	Domicilio bool		`bson:"domicilio" json:"domicilio"`
	Fecha_creacion time.Time `bson:"fecha_creacion" json:"fecha_creacion"`
	Estado int `bson:"estado" json:"estado"`
	Total	int	  `bson:"total" json:"total"`
}




