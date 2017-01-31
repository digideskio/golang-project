package models

import (  
    // Standard library packages
    "fmt"
	"time"

	// Third party packages
	_"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

    //Custom packages     
    "bitbucket.org/golang-project/todova_go_service/app/helpers"
    "bitbucket.org/golang-project/todova_go_service/databases"
)

type (
	// User represents the structure of our resource
	User struct {
		Id             bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`		
		Email        User            `json:"details" bson:"details"`
		Password        User            `json:"details" bson:"details"`
		PasswordSalt        User            `json:"details" bson:"details"`
		CreatedTime          time.Time       `json:"createdTime"             bson:"createdTime"`
        UpdatedTime          time.Time       `json:"updatedTime"             bson:"updatedTime"`
	}
)

// *****************************************************************************
// Model Methods
// *****************************************************************************

func UserInsert(user User) (ser,error) {

	//Get databaseName
	keysJson      := helpers.GetConfigKeys()
	envVariable   := helpers.GetEnvVariable()
	databaseName,_:=keysJson.String(envVariable,"databaseName")

	//Get mongoSession
	mangoSession:=databases.GetMongoSession()
	sessionCopy := mangoSession.Copy()
	defer sessionCopy.Close()

	col:=sessionCopy.DB(databaseName).C("user")
	err:= col.Insert(user)

	if err!= nil {
		fmt.Println(err)
		return user,err
	}

	return user,nil
}


func UserFindOne(query bson.M, selectQuery bson.M) (User, error){

	//Get databaseName
	keysJson      := helpers.GetConfigKeys()
	envVariable   := helpers.GetEnvVariable()
	databaseName,_:=keysJson.String(envVariable,"databaseName")

	//Get mongoSession
	mangoSession:=databases.GetMongoSession()
	sessionCopy := mangoSession.Copy()
	defer sessionCopy.Close()

	col:=sessionCopy.DB(databaseName).C("user")

	user:= User{}

	err := col.Find(query).Select(selectQuery).One(&user)
	if err!= nil {		
		fmt.Println(err)
		return user,err
	}

	return user,nil
}
