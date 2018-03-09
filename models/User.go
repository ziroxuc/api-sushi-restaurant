package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Nombre string `bson:"nombre,omitempty" json:"nombre,omitempty"`
	Usuario string `bson:"usuario" json:"usuario"`
	Password string `bson:"password" json:"password"`
	PasswordHash []byte `bson:"passwordHash,omitempty" json:"passwordHash,omitempty"`
	Email string	`bson:"email,omitempty" json:"email,omitempty"`
	Img string		`bson:"img,omitempty" json:"img,omitempty"`
	Role string 	`bson:"role,omitempty" json:"role,omitempty"`
	Create_date time.Time `bson:"create_date,omitempty" json:"create_date,omitempty"`
	Mod_date time.Time `bson:"mod_date,omitempty" json:"mod_date,omitempty"`
	Estado int `bson:"estado,omitempty" json:"estado,omitempty"`
	Token string 	`bson:"token,omitempty" json:"token,omitempty"`
}
