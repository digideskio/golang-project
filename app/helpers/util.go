package helpers

import (
	//Standard library packages
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"encoding/base64"

	//Third party packages
	"github.com/jmoiron/jsonq"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetConfigKeys()  *jsonq.JsonQuery{

	configPath,_:=filepath.Abs("config/keys.json")

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
	}
	var dat map[string] interface{}
	json.Unmarshal(file, &dat)

	jq :=jsonq.NewQuery(dat)
	return jq
}

func IsProduction() bool {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		return  false
	}

	return true
}

func GetEnvVariable() string {

	isProd:=IsProduction()
	if isProd {
		return "production"
	}

	return "development"
}

func GetPortAddress() string {

	var PORT string = ":3000" // If not found in env

	envport := os.Getenv("PORT")
	if envport != "" {
		PORT = ":" +os.Getenv("PORT")
	}

	return PORT
}

func GetAWSConfig(region string) (*aws.Config,error) {

	creds := credentials.NewEnvCredentials()

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad AWS credentials")
		return nil,err
	}

	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	return cfg,nil
}
