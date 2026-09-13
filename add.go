package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(u []Header) {

	// Taking inputs for struct var
	var tit, cont string

	randId := rand.Intn(9000) + 1000

	// Ensuring ID's uniqueness
restart1:
	for _, val := range u {

		if randId == val.Id {

			randId = rand.Intn(9000) + 1000
			goto restart1
		}
	}

	// Ensuring Title & Content has some value
restart2:
	fmt.Print("\nTitle : ")
	fmt.Scan(&tit)

	if tit == "" {

		fmt.Println("\nTitle cannot be empty. Please enter a valid Title....")
		goto restart2
	}

restart3:
	fmt.Print("\nContent : ")
	fmt.Scan(&cont)

	if cont == "" {

		fmt.Println("\nContent cannot be empty. Please enter some content....")
		goto restart3
	}

	ctime := time.Now().Format("01-02-2006 15:04:05 Monday")

	// Writing data to a Header variable
	user := Header{

		Id:         randId,
		Title:      tit,
		Content:    cont,
		Created_at: ctime,
		Updated_at: ctime,
	}

	// Storing Header var 'user' as a value of []Header slice 's1'
	s1 := []Header{user}

	// Appending s1 to existing slice data
	s := vomit()
	s = append(s, s1...)

	// Writing data to the file
	feed(s)

	// Printing success message
	fmt.Println("\n\nNOTE Created Successfully at ", ctime)
	fmt.Println("Id ~ ", randId)
	fmt.Println("Title ~ ", tit)

	// Asking whether user wants to add another note
	var ch string

	for {

		fmt.Print("\nDo you wish to add another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		// Redirecting to add() throught distro()
		case "y", "Y":
			distro(vomit(), "add")
			return

		// Redirecting to distro()
		case "n", "N":
			distro(vomit(), menu())
			return

		// Error; taking input again
		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}
}
