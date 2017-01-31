package validations

import (
	//Standard library packages
	"net/http"
	"encoding/json"	
	"bytes"
	"io"
	"io/ioutil"

	// Third party packages		
	"github.com/asaskevich/govalidator"

	//Custom packages
	"bitbucket.org/rtbathula/golang-project/app/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)

	decoder := json.NewDecoder(reader)

	var user models.User
	err := decoder.Decode(&user)

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

	if (user.Email == "") {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if !govalidator.IsEmail(user.Email) {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is invalid"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (user.Password == "") {
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

func LoginUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)

	decoder := json.NewDecoder(reader)

	var user models.User
	err := decoder.Decode(&user)

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

	if (user.Email == "") {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is required"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if !govalidator.IsEmail(user.Email) {
		var response models.Response
		response.Status  = "error"
		response.Message = "Email is invalid"
		respByt,_:= json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respByt)
		return
	}

	if (user.Password == "") {
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
