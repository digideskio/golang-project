package helpers

import (
    // Standard library packages  
    "fmt"  
    "os"
    "io"   
    "strings"
    "crypto/md5"
    "crypto/rand"      
    "encoding/base32"

    //Third party packages
    //Custom packages       
)

func ValidatePassword(reqPassword string, dbPasswordSalt string, dbPasswordHash string) bool {

    combination := dbPasswordSalt + string(reqPassword)
    reqPasswordHash := md5.New()
    io.WriteString(reqPasswordHash, combination)
    
    base32PwdStr := strings.ToLower(base32.HexEncoding.EncodeToString(reqPasswordHash.Sum(nil)))
     
    if(base32PwdStr==dbPasswordHash){
        return true
    } 

    return false     
}

func EncryptPassword(password string) (string,string) {   

    passwordByt := []byte(password)
    salt := generatePwdSalt(passwordByt)
    base32SaltStr := strings.ToLower(base32.HexEncoding.EncodeToString(salt))

    combination := base32SaltStr + string(password)
    passwordHash := md5.New()
    io.WriteString(passwordHash, combination)   
   
    base32PwdStr := strings.ToLower(base32.HexEncoding.EncodeToString(passwordHash.Sum(nil)))

    return base32SaltStr, base32PwdStr
}

func generatePwdSalt(secret []byte) []byte {

    const saltSize = 16

    buf := make([]byte, saltSize, saltSize+md5.Size)
    _, err := io.ReadFull(rand.Reader, buf)

    if err != nil {
        fmt.Printf("random read failed: %v", err)
        os.Exit(1)
    }

    hash := md5.New()
    hash.Write(buf)
    hash.Write(secret)
   
    return hash.Sum(buf)
}

func GetEncryptedPassword(password string) (string,string) {

    passwordByt := []byte(password)
    salt := generatePwdSalt(passwordByt)
    base32SaltStr := strings.ToLower(base32.HexEncoding.EncodeToString(salt))

    combination := base32SaltStr + string(password)
    passwordHash := md5.New()
    io.WriteString(passwordHash, combination)

    base32PwdStr := strings.ToLower(base32.HexEncoding.EncodeToString(passwordHash.Sum(nil)))

    return base32SaltStr, base32PwdStr
}