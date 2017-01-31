package controllers

import (
	// Standard library packages
	"errors"

	// Third party packages
	"gopkg.in/mgo.v2/bson"	

	//Custom packages
	"bitbucket.org/rtbathula/golang-project/app/models"
	"bitbucket.org/rtbathula/golang-project/app/helpers"
	_"bitbucket.org/rtbathula/golang-project/app/microservices"	
)

// ****************************************************************************
// Controllers Logic
// *****************************************************************************

func RegisterUser(user models.User) (models.Response,error) {

	var response models.Response
	var err error

	//Check email already exist
	query:=bson.M{"email":user.Email}
	queryVisibleFields := bson.M{}

	_, err = models.UserFindOne(query,queryVisibleFields)
	if (err == nil) {
		response.Status  = "error"
		response.Message = "Email already exist"
		return response,errors.New("Email already exist")
	}

	//SET Parameters to user
	pwdSalt, pwdStr := helpers.EncryptPassword(user.Password)
	user.Password = pwdStr
	user.PasswordSalt = pwdSalt

	_,err = models.UserInsert(user)
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

	var response models.Response

	queryEmail := bson.M{"email":email}		
	queryVisibleFields := bson.M{}
	
	user,err := models.UserFindOne(queryEmail,queryVisibleFields)
	if(err!= nil){		
		response.Status  = "error"
		response.Message = err.Error()
		return response, err
	}	

	isValid:=helpers.ValidatePassword(password,user.PasswordSalt,user.Password)

	if(!isValid){		
		response.Status  = "error"
		response.Message = "Invalid Password!"
		return response, errors.New("Invalid Password!")
	}
	
	response.Status  = "success"
	response.Message = "logged in successfully!"
	response.Result.DataType = "string"

	/* Create the token */
	jwtToken,err:=helpers.MakeJwtToken(user.Id)
	response.Result.Data = jwtToken
	
	return response, err
}
