package converter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type Converter struct {
	l *log.Logger
}

func NewConverter(l *log.Logger) *Converter {
	return &Converter{l}
}

//this method takes in a json object payload and converts it to a Go struct
func (ct Converter) JsonToStruct(w http.ResponseWriter, r *http.Request) {
	var payload []interface{}
	var response []interface{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
	}
	//evens are the field name and odds are the field type
	for i, value := range payload {
		dataType := reflect.TypeOf(value).String()

		if i%2 == 0 {
			response = append(response, value)
		} else {
			response = append(response, dataType)
		}
	}
	// fmt.Println(output)
	ct.l.Println(response)
	json.NewEncoder(w).Encode(response)
}
