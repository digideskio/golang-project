package helpers

import (
    // Standard library packages 
    _"encoding/json"
    _"io/ioutil"
    "fmt"
    "time"
    _"os"
    _"reflect"
    _"strings"

    // Third party packages
    "gopkg.in/mgo.v2/bson"       
    _"github.com/jmoiron/jsonq" 
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/auth0/go-jwt-middleware" 

    //Custom packages       
)

func MakeJwtToken(userId bson.ObjectId) (string, error) {

    keysJson    := GetConfigKeys()
    envVariable := GetEnvVariable()
        
    jwtSecret,keyErr:=keysJson.String(envVariable,"jwtSecret")
    if keyErr != nil {
        fmt.Println(keyErr)
        return "",keyErr
    }   

    /* Create the token */
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId"   : userId,        
        "exp"      : time.Now().Add(time.Hour * 24).Unix(),
    })      

    /* Sign the token with our secret */
    signedTkn, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        fmt.Println(keyErr)
        return "",err
    }

    return signedTkn,nil
}

func JwtValidateMiddleware() *jwtmiddleware.JWTMiddleware {

    keysJson    := GetConfigKeys()
    envVariable := GetEnvVariable()
        
    jwtSecret,keyErr:=keysJson.String(envVariable,"jwtSecret")
    if keyErr != nil {
        panic(keyErr.Error())
        return nil
    }

    var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
          return []byte(jwtSecret), nil
        },
        // When set, the middleware verifies that tokens are signed with the specific signing algorithm
        // If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
        // Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
        SigningMethod: jwt.SigningMethodHS256,
    })

    return jwtMiddleware
}

