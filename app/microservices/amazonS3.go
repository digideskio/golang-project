package microservices

import (
    //Standard library packages
    "fmt"
    "net/http"
    "bytes"
    "log"

    // Third party packages
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/aws/awsutil"

    //Custom packages
    "bitbucket.org/golang-project/todova_go_service/app/helpers"
)

func AmazonS3FileUpload(configName string, fileName string, buffer []uint8) (string,error) {

    envVariable  := helpers.GetEnvVariable()
    keysJson     := helpers.GetConfigKeys()

    //Config Names
    region,_     :=keysJson.String(envVariable,configName,"region")
    bucketName,_ :=keysJson.String(envVariable,configName,"bucketName")
    savePath,_   :=keysJson.String(envVariable,configName,"savePath")


    cfg,err:=helpers.GetAWSConfig(region)
    if err != nil {
        return "",err
    }

    //Set up a new s3manager client
    uploader := s3manager.NewUploader(session.New(cfg))

    //read into buffer
    fileBytes := bytes.NewReader(buffer)
    fileType := http.DetectContentType(buffer)

    path := savePath+"/"+fileName

    result, err := uploader.Upload(&s3manager.UploadInput{
        Bucket        : aws.String(bucketName),
        Key           : aws.String(path),
        Body          : fileBytes,
        ContentType   : aws.String(fileType),
        ACL           : aws.String("public-read"),
    })

    if err != nil {
        log.Fatalln("Failed to upload", err)
        return "",err
    }

    fmt.Printf("response %s", awsutil.StringValue(result))

    return result.Location,nil
}
