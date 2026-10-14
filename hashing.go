package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
// 	password := "password123"
// 	hash := sha256.Sum256([]byte(password))
// 	hash512:=sha512.Sum512([]byte(password))
// 	fmt.Printf("%X\n", hash)
// fmt.Printf("%X", hash512)
// 	// fmt.Println(string(password))
// 	p2 := "password123"
// 	fmt.Printf("%X", sha256.Sum256([]byte(p2)))

//generating salt
password:=generateSalt()

//Hash the password with salt
hash:=hashPassword(password, salt)
//store the salt and password in database, just printing as of now
sa
}
func generateSalt()([]byte, error){
	salt:=make([]byte,16 )
	_, err:=io.ReadFull(rand.Reader, salt)
	if err!=nil{
		return nil, err

	}

	return salt, nil
}
func HashPassword(password string, salt []byte)string {
	saltedPassword:=append(salt, []byte(password)...)
	hash:=sha256.Sum256(saltedPassword)

	return base64.StdEncoding.EncodeToString(hash[:])
	
}
