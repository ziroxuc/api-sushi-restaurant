package models

import "time"

type Paginator struct {

	OpcionNumber int `bson:"opcionNumber,omitempty" json:"opcionNumber,omitempty"`
	ItemsPerPage int `bson:"itemsPerPage" json:"itemsPerPage"`
	Pagina int `bson:"pagina" json:"pagina"`
	TotalRows int `bson:"totalRows,omitempty" json:"totalRows,omitempty"`
	FechaDesde time.Time `bson:"fechadesde,omitempty" json:"fechadesde,omitempty"`
	FechaHasta time.Time `bson:"fechahasta,omitempty" json:"fechahasta,omitempty"`
}