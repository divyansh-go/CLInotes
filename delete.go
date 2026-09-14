package main

import (
	"fmt"
)

func delete(u []Header, ch int) {

	// Printing info of the chosen noteid
	fmt.Println("\nYou are about to delete:")
	view(u, ch)

	// Taking index value of the chosen noteid
	var index int

	for i, val := range u {

		if val.Id == ch {

			index = i
		}
	}

	// Confirming deletion
loop:

	for {

		fmt.Print("\nAre you sure? (Y/N): ")
		sch := get()

		switch sch {

		// Performing deletion & breaking the loop
		case "y", "Y":
			u = append(u[:index], u[index+1:]...)

			// Feeding updated data to the file
			feed(u)

			// Printing success message
			fmt.Println("\nNOTE Deleted Successfully.........")

			break loop

		// Redirecting to distro()
		case "n", "N":
			distro(u, menu())
			return

		// Error; taking another input
		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}

	// Asking whether user wants to delete another Note
	for {

		fmt.Print("\nDo you wish to delete another note? (Y/N): ")
		sch := get()

		switch sch {

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
