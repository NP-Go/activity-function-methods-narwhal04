package main

import "fmt"

func main() {
	//Declare new item customer
	customer1 := custStruct{
		fName:   "Micheal",
		lName:   "Jordan",
		uName:   "MJ2020",
		passwd:  "1234567",
		email:   "MJ2020@gmail.com",
		phone:   12345678,
		address: "18227 Capstan Greens Road Cornelius, NC 28031.",
	}

	//Insert Code here
	fmt.Println(customer1.userCredentials())
	fmt.Println(customer1.userAddress())
	fmt.Println(customer1.allUserInfor())

}
