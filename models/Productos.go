package models

import "gopkg.in/mgo.v2/bson"

type Producto struct {
	ID              bson.ObjectId `bson:"_id" json:"_id"`
	Id_sushi		int			  `bson:"id_sushi" json:"id_sushi"`
	Id_categoria	int			  `bson:"id_categoria" json:"id_categoria"`
	Nombre      	string        `bson:"nombre" json:"nombre"`
	Precio  		int        	  `bson:"precio" json:"precio"`
	Precio_antiguo  int        	  `bson:"precio_antiguo" json:"precio_antiguo"`
	Piezas      	int        	  `bson:"piezas" json:"piezas"`
	Descripcion 	string        `bson:"descripcion" json:"descripcion"`
	Img 			string        `bson:"img" json:"img"`
	Create_date 	string        `bson:"create_date" json:"create_date"`
	Mod_date 		string        `bson:"mod_date" json:"mod_date"`
	Estado 			int           `bson:"estado" json:"estado"`
	Cantidad 		int			  `bson:"cantidad" json:"cantidad"`
	Destacado		bool		  `bson:"destacado" json:"destacado"`
	Stock			bool		  `bson:"stock" json:"stock"`
	Personalizaciones				  `bson:"personalizacion" json:"personalizacion"`
}

type Personalizacion struct {
	Titulo string `bson:"titulo" json:"titulo"`
	Opciones 	  `bson:"opciones" json:"opciones"`
}

type Personalizaciones []Personalizacion

type Opcion struct {
	Nombre string 		`bson:"nombre" json:"nombre"`
	Precio int 			`bson:"precio" json:"precio"`
	Seleccionada bool 	`bson:"seleccionada" json:"seleccionada"`

}
type Opciones []Opcion

type Productos []Producto
