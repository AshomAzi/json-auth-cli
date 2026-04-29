package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
} 

func Register() {
	var firstName string
	fmt.Print("Input your first name, must be more than 2 letters: ")
start:
	fmt.Scanln(&firstName)
	if len(firstName) >= 3 {
		var lastName string
		fmt.Print("Input your last name, must be more than 2 letters: ")
	start1:
		fmt.Scanln(&lastName)
		if len(lastName) >= 3 {
			var username string
			fmt.Print("Choose a username, must be more than 4 characters: ")
		start2:
			fmt.Scanln(&username)
			if len(username) >= 5 {
				var password1 string
				fmt.Print("Choose a password, must be more than 5 characters: ")
			start3:
				fmt.Scanln(&password1)
				if len(password1) >= 6 {
					var password2 string
					fmt.Print("Confirm Password: ")
				start4:
					fmt.Scanln(&password2)
					if password1 == password2 {
						var user Users
						user.FirstName = firstName
						user.LastName = lastName
						user.Username = username
						user.Password = password1
						fmt.Printf("%v\n", user)
						credential, err := json.Marshal(user)
						if err != nil {
							fmt.Println("An error occured when creating user!!!")
						}
						log.Printf("%s\n", credential)
						file, err := os.OpenFile("data.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							fmt.Println("An error occured:", err)
						}
						defer file.Close()
						file.Write(append(credential, '\n'))
						fmt.Println("User created successfully")
					} else {
						fmt.Print("Password does not match!!! ")
						goto start4
					}
				} else {
					fmt.Print("Password must be more than 5 characters: ")
					goto start3
				}
			} else {
				fmt.Print("Username must be more thst 4 letters: ")
				goto start2
			}

		} else {
			fmt.Print("name must be more than 2 letters: ")
			goto start1
		}
	} else {
		fmt.Print("name must be more than 2 letters: ")
		goto start
	}
}

func SignUp() {
	var username string
	fmt.Print("Input you user name: ")
	fmt.Scanln(&username)
	file, err := os.ReadFile("data.jsonl")
	check(err)
	for _, v := range string(file){
		if string(v) == "123456" {
			fmt.Println(v)
		}
	}
}

func main() {
	var command string
	fmt.Print("Welcome, Choose \"R\" to register and \"L\" to login: ")
	fmt.Scanln(&command)
	if command == "R" || command == "r" {
		Register()
	} else if command == "L" || command == "l" {
		SignUp()
	} else {
		fmt.Println("Invalid Command!!!")
	}
}
