package controllers

import (
	// Standard library packages
	"errors"
	_"net/url"
	_"fmt"

	// Third party packages
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	_"github.com/mattbaird/gochimp"
	_"github.com/rs/xid"

	//Custom packages
	"bitbucket.org/golang-project/todova_go_service/app/models"
	"bitbucket.org/golang-project/todova_go_service/app/helpers"
	_"bitbucket.org/golang-project/todova_go_service/app/microservices"	
)

// ****************************************************************************
// Controllers Logic
// *****************************************************************************

type AuthAdminResp struct {
	Admin models.Admin 	`json:"admin"`
	JwtToken string		`json:"jwtToken"`
}

func RegisterUser(user models.User) (models.Response,error) {

	var response models.Response

	//Check email already exist
	query:=bson.M{"email":admin.Details.Email}
	queryVisibleFields := bson.M{}

	_, err := models.UserFindOne(query,queryVisibleFields)
	if (err == nil) {
		response.Status  = "error"
		response.Message = "Email already exist"
		return response,errors.New("Email already exist")
	}

	//SET Parameters to user
	pwdSalt, pwdStr := helpers.EncryptPassword(admin.Details.Password)
	user.Password = pwdStr
	user.PasswordSalt = pwdSalt

	user_insert,err := models.UserInsert(user)
	if(err!= nil){
		response.Status  = "error"
		response.Message = "failed to register user"
		return response,err
	}

	//Continue
	response.Status  = "success"
	response.Message = "successfully registered!"	
	return response,nil
}

func LoginUser(email string, password string)(models.Response,error){

	queryEmail := bson.M{"details.email":email,"details.userType": "admin"}
	
	var queryVisibleFields bson.M
	queryVisibleFields= make(map[string]interface {})
	queryVisibleFields["details.createdTime"] = 0
	queryVisibleFields["details.emailConfirmationToken"] = 0
	queryVisibleFields["details.createdTime"] = 0
	queryVisibleFields["details.facebookAccessToken"] = 0
	queryVisibleFields["details.updatedTime"] = 0
	queryVisibleFields["details.resetPasswordToken"] = 0	
	
	admin,err := models.AdminFindOne(queryEmail,queryVisibleFields)

	if(err!= nil){
		var response models.Response
		response.Status  = "error"
		response.Message = err.Error()
		return response, err
	}	

	isValid:=helpers.ValidatePassword(password,admin.Details.PasswordSalt,admin.Details.Password)

	if(!isValid){
		var response models.Response
		response.Status  = "error"
		response.Message = "Invalid Password!"
		return response, err
	}		

	var response models.Response
	response.Status  = "success"
	response.Message = "login successfully!"
	response.Result.DataType = "json"

	/* Create the token */
	jwtToken,err:=helpers.MakeJwtToken(admin.Id)

	loginResponse := AuthAdminResp{}
	loginResponse.Admin = admin
	loginResponse.JwtToken = jwtToken

	response.Result.Data = loginResponse
	return response, err
}
