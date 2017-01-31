package main

import (
	//Standard library packages
	"net/http"
	"fmt"

	//Third party packages
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	//Custom packages
	"bitbucket.org/rtbathula/golang-project/app/api"
	"bitbucket.org/rtbathula/golang-project/app/helpers"
	"bitbucket.org/rtbathula/golang-project/databases"	
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("Golang server is up and running successfully!"))
	})

	n := negroni.Classic() // Includes some default middlewares
	n.Use(cors.New(cors.Options{
		AllowedOrigins   : []string{"*"},
		AllowedMethods   : []string{"GET","POST","DELETE","PUT", "PATCH"},
		AllowedHeaders   : []string{"Origin","Authorization"},
		ExposedHeaders   : []string{"Content-Length"},
		AllowCredentials : true,
	}))
	n.UseHandler(r)

	//Connect MongoDB
	initMongoDB();

	//Routes
	routes(r)

	//Run server
	http.Handle("/", r)
	http.ListenAndServe(helpers.GetPortAddress(), n)
}

//Private Fuctions
func initMongoDB(){
	databases.ConnectDB()
}


func routes(r *mux.Router){	
	api.UserApi(r)
}
