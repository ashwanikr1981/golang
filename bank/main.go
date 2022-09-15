package main

import (
	"fmt"
	"os"
)

type User struct {
	Name     string
	Id       int
	Password string
	Balance  int64
}

var Account []User

func login() {

	var userid int
	var password string
START:
	fmt.Print("Please enter your user id: ")
	fmt.Scan(&userid)
	fmt.Print("Please enter your password: ")
	fmt.Scan(&password)

	var isFound bool = false

	for _, val := range Account {
		if userid == val.Id && password == val.Password {
			afterLoginOptions()
		} else {
			fmt.Println("out")
			isFound = false
		}
	}

	if !isFound {
		fmt.Println("Invalid Credentials. Type c for try again and q for quit")
		var opt string
		fmt.Print("Enter: ")
		fmt.Scan(&opt)
		if opt == "c" {
			goto START
		} else {
			fmt.Println("Bye!!!")
			exitProgram()
		}
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
	case 2:
		depositMoney()
	case 3:
		fmt.Println("balance")
		fmt.Println("Bye!!")
		exitProgram()
	case 4:
		fmt.Println("Bye!!")
		exitProgram()
	default:
		fmt.Println("Invalid please try again")
		afterLoginOptions()
	}

}

func exitProgram() {
	os.Exit(1)
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
	fmt.Println("Create Account")
	fmt.Print("Please enter name: ")
	fmt.Scan(&name)
	fmt.Print("Please enter user id: ")
	fmt.Scan(&userid)
	fmt.Print("Please enter password: ")
	fmt.Scan(&password)

	userData := User{Name: name, Id: userid, Password: password}

	Account = append(Account, userData)

	fmt.Println("User has been created successfully", Account)
	main()
}

func main() {
	fmt.Println("Hi! Welcome to Mr. Ashwani ATM Machine!")
	fmt.Println("Please select an option from the menu below:")
	fmt.Println("l -> Login\nc -> Create New Account\nq -> Quit")
	fmt.Print("Enter: ")
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
		exitProgram()
	} else {
		fmt.Println("Wrong selection. Please try again!")
		main()
	}

}
