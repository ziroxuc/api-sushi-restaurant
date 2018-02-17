package models

type Producto struct {
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
	Personalizacion
}

type Personalizacion struct {
	Titulo string `bson:"titulo" json:"titulo"`
	Opciones
}

type Opciones struct {
	Nombre string 		`bson:"nombre" json:"nombre"`
	Precio int 			`bson:"precio" json:"precio"`
	Seleccionado bool 	`bson:"seleccionado" json:"seleccionado"`

}

type Productos []Producto
