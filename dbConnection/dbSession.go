package dbConnection

import "gopkg.in/mgo.v2"

const SERVER = "mongodb://admin:admin123@localhost:27017/"
const DATABASE_NAME = "sushi-restaurant-db"
const USER  = "admin"
const PASSWORD  = "admin123"


func GetSession() *mgo.Session {
	session, error := mgo.Dial(SERVER)
	if (error != nil) {
		panic(error)
	}
	return session
}

func GetCollectionPedidos() *mgo.Collection{
	var collectionPedidos = GetSession().DB(DATABASE_NAME).C("pedidos")
	return collectionPedidos
}

func GetCollectionProductos() *mgo.Collection{
	var collectionProductos= GetSession().DB(DATABASE_NAME).C("productos")
	return collectionProductos
}

func GetCollectionEstados() *mgo.Collection{
	var collectionProductos= GetSession().DB(DATABASE_NAME).C("estados")
	return collectionProductos
}

func GetCollectionCategorias() *mgo.Collection{
	var collectionProductos= GetSession().DB(DATABASE_NAME).C("categorias")
	return collectionProductos
}
func GetCollectionPromociones() *mgo.Collection{
	var collectionProductos= GetSession().DB(DATABASE_NAME).C("promociones")
	return collectionProductos
}
func GetCollectionUsuario() *mgo.Collection{
	var collectionUsuario = GetSession().DB(DATABASE_NAME).C("usuario")
	return collectionUsuario
}

