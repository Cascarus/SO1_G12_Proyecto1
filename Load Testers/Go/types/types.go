package types

type Tuit struct {

	Nombre string `json:"nombre"`
	Comentario string `json:"comentario"`
	Fecha string `json:"fecha"`
	Hashtags []string `json:"hashtags"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`

}

type Message struct{
	Data string
}

type Log struct {
    StatusNumber int `json:"statusNumber"`
    Message string `json:"message"`
}