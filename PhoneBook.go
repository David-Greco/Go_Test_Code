package main

import "fmt"

//David Greco
//Basic Phone Book Implementation
//Add user, remove, search
//No functions used

type user struct {
	name   string
	number int
}

type phoneBook []user

func main() {

	myBook := make(phoneBook, 0)

	//Add user
	myBook = append(myBook, user{"David Greco", 6095551234})
	fmt.Println(myBook)

	//Search user
	DavidFound := false
	for x := range myBook {
		if myBook[x].name == "David Greco" {
			DavidFound = true
			fmt.Println("Found David")
		}
	}
	if DavidFound == false {
		fmt.Println("David Not found")
	}

	//Delete user
	index := 0
	for x := range myBook {
		if myBook[x].name == "David Greco" {
			index = x
		}
	}
	myBook = append(myBook[:index], myBook[index+1:]...)

	//Search deleted user
	DavidFound = false
	for x := range myBook {
		if myBook[x].name == "David Greco" {
			DavidFound = true
			fmt.Println("Found David")
		}
	}
	if DavidFound == false {
		fmt.Println("David Not found")
	}

}
