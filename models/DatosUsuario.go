package models

type DatosUsuario struct {
	Nombre  	string        	  	`bson:"nombre" json:"nombre"`
	Email      	string        	  	`bson:"email" json:"email"`
	Retiro 		string        		`bson:"retiro" json:"retiro"`
	Telefono	string        		`bson:"telefono" json:"telefono"`
	Direccion   string        		`bson:"direccion" json:"direccion"`
	Comuna		string			  	`bson:"comuna" json:"comuna"`
}

type DatosUsuarios []DatosUsuario
