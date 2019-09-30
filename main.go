// Author: Mathis Van Eetvelde 			Mathisco-01

package main 

import (
	"fmt"
	"./hashing"
)

func main(){


	var username string
	var password string

	//Input "username" and "password"
	fmt.Print("Type username: ")
	fmt.Scan(&username)
	fmt.Print("Type password: ")
	fmt.Scan(&password)

	hashedLogin := hashing.HashLoginInfo(username, password)

	fmt.Println("Hashed username: " + hashedLogin.Hashed_username)
	fmt.Println("Hashed password: " + hashedLogin.Hashed_password)
	fmt.Println("Hashing salt: " + hashedLogin.Password_salt)
}