package api

import (
	// Standard library packages
	_"fmt"
	"net/http"
	"encoding/json"

	// Third party packages
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	//Custom packages
	"bitbucket.org/golang-project/todova_go_service/app/controllers"
	"bitbucket.org/golang-project/todova_go_service/app/models"
	"bitbucket.org/golang-project/todova_go_service/app/helpers"
)

// *****************************************************************************
// API Routes
// *****************************************************************************

func UserApi(r *mux.Router) {

	r.Handle("/user/register",negroni.New(	
		negroni.HandlerFunc(validations.RegisterUser),	
		negroni.Wrap(RegisterUser),
	)).Methods("POST")

	r.Handle("/user/login",negroni.New(
		negroni.HandlerFunc(validations.LoginUser),	
		negroni.Wrap(LoginUser),
	)).Methods("POST")
}


/*Desc   : Register User
  Params : model User{email,password}
  Returns: Success message or error
*/
var RegisterUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user models.User

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)

	resp, err := controllers.RegisterUser(user)
	respByt,_:= json.Marshal(resp)

	if err != nil {		
		w.WriteHeader(400)
		w.Write(respByt)
		return
	}
	
	w.WriteHeader(200)
	w.Write(respByt)
	return
})

/*Desc   : Login Admin
  Params : model User{email,password}
  Returns: User or error
*/
var LoginUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
	decoder := json.NewDecoder(r.Body)

	var user models.User
	err := decoder.Decode(&user)

	resp, err := controllers.LoginUser(user.Email, user.Password)
	respByt,_:= json.Marshal(resp)

	if err != nil {
		w.WriteHeader(400)
		w.Write(respByt)
		return
	}	

	w.WriteHeader(200)
	w.Write(respByt)
	return
})
