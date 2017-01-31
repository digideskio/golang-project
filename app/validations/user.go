package validations

import (
	//Standard library packages
	"net/http"
	"encoding/json"	
	"bytes"
	"io"
	"io/ioutil"

	// Third party packages
	"github.com/gorilla/mux"
	"github.com/crowl/rut"
	"github.com/asaskevich/govalidator"

	//Custom packages
	"bitbucket.org/golang-project/todova_go_service/app/models"
)

func RegisterCustomer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)

	decoder := json.NewDecoder(reader)

	customer := models.Customer{}
	err := decoder.Decode(&customer)

	r.Body = ioutil.NopCloser(b)

	if err != nil {
		var response models.Response
		response.Status  = "error"
		response.Message = "Invalid paramss"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (customer.Email == "") {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if !govalidator.IsEmail(customer.Email) {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is invalid"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (customer.Password == "") {
		var response models.Response
		response.Status  = "error"
		response.Message = "Password is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	next(w,r)
}

func LoginCustomer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)

	type Request struct {
		Email				string	`json:"email"`
		Password			string	`json:"password"`		
	}

	decoder := json.NewDecoder(reader)

	req := Request{}
	err := decoder.Decode(&req)

	r.Body = ioutil.NopCloser(b)

	var response models.Response
	if err != nil {		
		response.Status  = "error"
		response.Message = "Parámetros inválidos!"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if req.Email == "" {	
		response.Status  = "error"
		response.Message = "Email is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if req.Password == "" {
		
		response.Status  = "error"
		response.Message = "Password is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	next(w, r)
}

