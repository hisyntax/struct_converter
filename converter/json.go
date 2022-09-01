package converter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Converter struct {
	l *log.Logger
}

func NewConverter(l *log.Logger) *Converter {
	return &Converter{l}
}

//this method takes in a json object payload and converts it to a Go struct
func (ct Converter) JsonToStruct(w http.ResponseWriter, r *http.Request) {
	var payload interface{}
	// var output
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		fmt.Println(err)
	}

	ct.l.Println(payload)
}
