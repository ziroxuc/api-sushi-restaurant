package dbConnection

import "gopkg.in/mgo.v2"

const SERVER = "mongodb://localhost"
const DATABASE_NAME = "sushi-restaurant-db"

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