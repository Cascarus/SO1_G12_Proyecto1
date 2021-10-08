package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tuit struct {

    ID primitive.ObjectID `json:"_id,omitempty"`
	Nombre string `json:"nombre"`
	Comentario string `json:"comentario"`
	Fecha string `json:"fecha"`
	Hashtags []string `json:"hashtags"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`

}

type Log struct {
    StatusNumber int `json:"statusNumber"`
    Message string `json:"message"`
	Time time.Duration `json:"time"`
}


type Message struct {
	Guardados int `json:"guardados"`
	Api string `json:"api"`
	TiempoCarga string `json:"tiempoCarga"`
	Db string `json:"db"`
}