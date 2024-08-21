package productcontroller

import (
	"encoding/json"
	"fmt"
	"go-starter-webapp/app/models/productmodel"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	data := productmodel.Getall()
	b, _ := json.Marshal(data)
	fmt.Fprint(writer, string(b))
}
func Create(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "test")
}
