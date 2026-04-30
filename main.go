package main

import (
	"cli/database"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Register() {
	firstName := bufio.NewScanner(os.Stdin)
	fmt.Print("Input your first name, spaces are not allowed, must be more than 2 letters: ")
start:
	firstName.Scan()
	fName := firstName.Text()
	firstFullName := strings.Split(fName, " ")
	if len(firstFullName) == 1 {
		FirstName := firstFullName[0]
		if len(FirstName) > 2 {
			lastName := bufio.NewScanner(os.Stdin)
			fmt.Print("Input your last name, spaces are not allowed, must be more than 2 letters: ")
		start1:
			lastName.Scan()
			lName := lastName.Text()
			lastFullName := strings.Split(lName, " ")
			if len(lastFullName) == 1 {
				LastName := lastFullName[0]
				if len(LastName) > 2 {
					userName := bufio.NewScanner(os.Stdin)
					fmt.Print("Choose a username, spaces are not allowed. Must be more than 5 characters without spaces: ")
				start2:
					userName.Scan()
					FullUserName := userName.Text()
					SingleUserName := strings.Split(FullUserName, " ")
					if len(SingleUserName) == 1 {
						UserName := SingleUserName[0]
						if len(UserName) > 5 {
							fmt.Print("Choose a password, must be more than 5	 characters: ")
						start3:
							password1 := bufio.NewScanner(os.Stdin)
							password1.Scan()
							Pass1 := password1.Text()
							if len(Pass1) >= 6 {
								fmt.Print("Confirm Password: ")
							start4:
								password2 := bufio.NewScanner(os.Stdin)
								password2.Scan()
								Pass2 := password2.Text()
								if Pass1 == Pass2 {
									var user database.Users
									user.FirstName = FirstName
									user.LastName = LastName
									user.UserName = UserName
									user.Password = Pass1
									data, err := json.Marshal(user)
									if err != nil {
										fmt.Println("An error occured when creating user!!!")
									}
									file, err := os.OpenFile("user_credentials.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
									if err != nil {
										fmt.Println("An error occured:", err)
									}
									defer file.Close()
									file.Write(append(data, '\n'))
									fmt.Println("User created successfully")
								} else {
									fmt.Print("Password does not match. Try again: ")
									goto start4
								}
							} else {
								fmt.Print("Password must be more than 5 characters: ")
								goto start3
							}

						} else {
							fmt.Print("Username must be more than 5 characters: ")
							goto start2
						}

					} else {
						fmt.Print("Username must not contains any spaces:  ")
						goto start2
					}

				} else {
					fmt.Print("Lastname must be more than 2 letters. Try again: ")
					goto start1
				}

			} else {
				fmt.Print("Input Just your lastname, spaces are not allowed. Try again: ")
				goto start1
			}

		} else {
			fmt.Print("Firstname must be more than 2 letters, spaces are not allowed. Try again: ")
			goto start
		}

	} else {
		fmt.Print("Input just your first name, spaces are not allowed. Try again: ")
		goto start
	}
}

func SignUp() {
	var username string
	fmt.Print("Input you user name: ")
	fmt.Scanln(&username)
	file, err := os.ReadFile("data.jsonl")
	check(err)
	for _, v := range string(file) {
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
