package main

import (
	"fmt"
	"strings"
)

func delete(u []Header, ch1 int) {

	// Printing info of the chosen noteid
	fmt.Println("\nYou are about to delete:")
	view(u, ch1)

	// Taking index value of the chosen noteid
	var index int

	for i, val := range u {

		if val.Id == ch1 {

			index = i
		}
	}

	// Confirming deletion
loop:

	for {

		var sch string

		fmt.Print("\nAre you sure? (y/n): ")
		fmt.Scan(&sch)
		lsch := strings.ToLower(sch)

		switch lsch {

		// Performing deletion & breaking the loop
		case "y":
			u = append(u[:index], u[index+1:]...)

			// Feeding updated data to the file
			feed(u)

			// Printing success message
			fmt.Println("\nNOTE Deleted Successfully.........")

			break loop

		// Redirecting to distro()
		case "n":
			distro(u, menu())
			return

		// Error; taking another input
		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}

	// Asking whether user wants to delete another Note
	var ch string

	for {

		fmt.Print("\nDo you wish to delete another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		// Redirecting to delete() through distro()
		case "y", "Y":
			distro(nil, "delete")
			return

		// Redirecting to distro()
		case "n", "N":
			distro(vomit(), menu())
			return

		// Error; taking another input
		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}
}
