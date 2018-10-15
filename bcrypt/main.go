package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	paswd := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(paswd), bcrypt.MinCost)

	if err != nil {
		fmt.Println("Error while hashing password: ", err)
	}

	fmt.Println("Stored Password:", paswd)
	fmt.Println("Hash of the password:", string(bs))

	//Read password from user
	fmt.Println("Enter the password")
	var loginPwd string

	fmt.Scan(&loginPwd)

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("Invalid Password")
		return
	}
	fmt.Println("You're logged In!")
}
