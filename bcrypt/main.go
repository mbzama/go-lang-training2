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

	fmt.Println(paswd)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPwd := "password1230"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("Invalid Password")
		return
	}
	fmt.Println("You're logged In!")
}
