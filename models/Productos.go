package models

type Producto struct {
	Id_sushi	int			  `bson:"id_sushi" json:"id_sushi"`
	Nombre      string        `bson:"nombre" json:"nombre"`
	Precio  	int        	  `bson:"precio" json:"precio"`
	Piezas      int        	  `bson:"piezas" json:"piezas"`
	Descripcion string        `bson:"descripcion" json:"descripcion"`
	Img 		string        `bson:"img" json:"img"`
	Create_date string        `bson:"create_date" json:"create_date"`
	Mod_date 	string        `bson:"mod_date" json:"mod_date"`
	Estado 		int           `bson:"estado" json:"estado"`
	Cantidad 	int			  `bson:"cantidad" json:"cantidad"`
}

type Productos []Producto
