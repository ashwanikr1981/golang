package main

import (
	"fmt"
	"os"
)

type User struct {
	Name     string
	Id       int
	Password string
}

var UserId int = 11
var UserPassword string = "1234"

func login() {
	var userid int
	var password string
	fmt.Print("Please enter your user id: ")
	fmt.Scan(&userid)
	fmt.Print("Please enter your password: ")
	fmt.Scan(&password)

	if userid == UserId && password == UserPassword {
		afterLoginOptions()
	} else {
		fmt.Println("Invalid Credentials. Please try again")
		main()
	}
}

func afterLoginOptions() {
	fmt.Println("Please select an option from the menu below:")
	fmt.Println("1 -> Withdraw Money")
	fmt.Println("2 -> Deposit Money")
	fmt.Println("3 -> Request Balance")
	fmt.Println("4 -> Quit the Program")

	var opt int

	fmt.Scan(&opt)

	switch opt {
	case 1:
		withDrawMoney()
		os.Exit(1)
	case 2:
		depositMoney()
		os.Exit(1)
	case 3:
		fmt.Println("balance")
		fmt.Println("Bye!!")
		os.Exit(1)
	case 4:
		fmt.Println("Bye!!")
		os.Exit(1)
	default:
		fmt.Println("Invalid please try again")
		afterLoginOptions()
	}

}

func withDrawMoney() {
	fmt.Print("Enter amount to be withdrawn: ")
	var amount int64
	fmt.Scan(&amount)
	fmt.Println("Thanks for withdrawing amount: ", amount)
	afterLoginOptions()
}

func depositMoney() {
	fmt.Print("Enter amount to be deposited: ")
	var amount int64
	fmt.Scan(&amount)
	fmt.Println("Thanks for depositing amount: ", amount)
	afterLoginOptions()
}

func createAccount() {
	var userid int
	var name, password string
	fmt.Print("Please enter name: ")
	fmt.Scan(&name)
	fmt.Print("Please enter user id: ")
	fmt.Scan(&userid)
	fmt.Print("Please enter password: ")
	fmt.Scan(&password)

	userData := User{name, userid, password}

	fmt.Println("User has been created successfully", userData)
	main()
}

func main() {
	fmt.Println("Hi! Welcome to Mr. Ashwani ATM Machine!")
	fmt.Println("Please select an option from the menu below:")
	fmt.Println("l -> Login\nc -> Create New Account\nq -> Quit")

	var opt string

	_, err := fmt.Scanln(&opt)
	if err != nil {
		panic(err)
	}

	if opt == "l" {
		login()
	} else if opt == "c" {
		createAccount()
	} else if opt == "q" {
		fmt.Println("Thank you for using our service. Bye!")
	} else {
		fmt.Println("Wrong selection. Please try again!")
		main()
	}

}
