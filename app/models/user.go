package models

import (  
    // Standard library packages
    "fmt"
	"time"

	// Third party packages
	_"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

    //Custom packages     
    "bitbucket.org/rtbathula/golang-project/app/helpers"
    "bitbucket.org/rtbathula/golang-project/databases"
)

type (
	// User represents the structure of our resource
	User struct {
		Id             bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`		
		Email          string            `json:"email" bson:"email"`
		Password       string            `json:"password" bson:"password"`
		PasswordSalt   string            `json:"passwordSalt" bson:"passwordSalt"`
		CreatedTime    time.Time         `json:"createdTime"  bson:"createdTime"`
        UpdatedTime    time.Time         `json:"updatedTime"  bson:"updatedTime"`
	}
)

// *****************************************************************************
// Model Methods
// *****************************************************************************

func UserInsert(user User) (User,error) {

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
