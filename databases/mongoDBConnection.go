package databases

import (
    //Standard library packages

    //Third party packages  
    "gopkg.in/mgo.v2"

    //Custom packages  
    "bitbucket.org/CarlaRod/todova_go_service/app/helpers"

)

var mongoSession *mgo.Session

func ConnectDB(){

    keysJson    := helpers.GetConfigKeys()
    envVariable := helpers.GetEnvVariable()

    dbConnection, keyErr := keysJson.String(envVariable, "mongoDBConnection")
    if keyErr != nil {
        panic(keyErr)
    }

    //Connect to our mongo
    s, err := mgo.Dial(dbConnection)
    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    mongoSession =s
}

func GetMongoSession() *mgo.Session{
    return mongoSession
}

func CreateIndexes(collectionName string, params map[string]string)  {

    //Get databaseName
    keysJson        := helpers.GetConfigKeys()
    envVariable     := helpers.GetEnvVariable()
    databaseName, _ := keysJson.String(envVariable,"databaseName")

    //Get mongoSession
    sessionCopy := mongoSession.Copy()
    defer sessionCopy.Close()

    col := sessionCopy.DB(databaseName).C(collectionName)

    keys := []string{}

    for key, value := range params {
        keys = append(keys, key+":"+value)
    }

    index := mgo.Index{
        Key: keys,
    }

    err := col.EnsureIndex(index)
    if err != nil {
        panic(err)
    }
}


//Local: mongodb://localhost