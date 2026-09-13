package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("==================================")
	fmt.Println("\n             NOTEBOOK             ")
	fmt.Println("\n==================================")

	distro(vomit(), menu())

	fmt.Println("\n\nThank you for using Notebook")
	fmt.Println("\nGoodbye!")
}

func menu() string {

	fmt.Println("\n\n------MENU------")
	var ch string

	fmt.Println("\n Options Available:\n ~ Add\n ~ List\n ~ View\n ~ Update\n ~ Delete\n ~ Exit")
	fmt.Print("\n\nChoose an action to perform: ")
	fmt.Scan(&ch)

	return ch
}

func distro(u []Header, str string) {

	lch := strings.ToLower(str)

	switch lch {

	case "add":
		fmt.Println("\n\n------ADD NOTES------")
		add(u)

	case "list":
		fmt.Println("\n\n------ALL NOTES------")
		list(u)
		recurrlist()

	case "view":
		fmt.Println("\n\n------VIEW NOTES------")
		list(u)
		view(u, ask())
		recurrview()

	case "update":
		fmt.Println("\n\n------UPDATE NOTES------")
		list(u)
		update(u, ask())

	case "delete":
		fmt.Println("\n\n------DELETE NOTES------")
		list(u)
		delete(u, ask())

	case "exit":
		return

	default:
		fmt.Println("\nInvalid Action. Please choose a valid action to perform......")
		distro(vomit(), menu())

	}

}

// Takes input for Note ID
func ask() int {

	for {

		var ch int
		fmt.Print("\nEnter note ID: ")
		fmt.Scan(&ch)

		// Checking whether the given Note ID exists
		for _, val := range vomit() {

			if ch == val.Id {

				return ch
			}
		}

		fmt.Println("Note ID not found. Please try again.......")
	}
}
