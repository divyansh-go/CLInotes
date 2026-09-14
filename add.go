package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(u []Header) {

	randId := rand.Intn(9000) + 1000

	// Ensuring ID's uniqueness
restart:
	for _, val := range u {

		if randId == val.Id {

			randId = rand.Intn(9000) + 1000
			goto restart
		}
	}

	// Taking inputs for Title & Content
	var tit, cont string

	for {

		fmt.Print("\nTitle : ")
		tit = get()

		// Ensuring Title has some value
		if tit != "" {

			break
		}

		fmt.Println("\nTitle cannot be empty. Please enter a valid Title....")
	}

	for {

		fmt.Println("\nType '#' and Press 'Enter' to stop writing.")

		fmt.Print("\nContent : ")
		cont = getlines()

		// Ensuring Content has some value
		if cont != "" {

			break
		}

		fmt.Println("\nContent cannot be empty. Please enter some content....")
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
	for {

		fmt.Print("\nDo you wish to add another note? (Y/N): ")
		ch := get()

		switch ch {

		// Redirecting to add() through distro()
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
