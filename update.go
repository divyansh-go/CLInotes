package main

import (
	"fmt"
	"strings"
	"time"
)

func update(u []Header, ch1 int) {

	// Viewing given Header var before deletion
	view(u, ch1)

	// Getting index value of the given Note ID
	var index int

	for i, val := range u {

		if val.Id == ch1 {

			index = i
		}
	}

	for {

		c := 0

		// Taking the field to be updated
		var sch string

		fmt.Print("\nEnter the field to be updated. (Title/Content/Both): ")
		fmt.Scan(&sch)
		lsch := strings.ToLower(sch)

		// Updating title and/or content
		if lsch == "title" || lsch == "both" {

			fmt.Print("\nEnter New Title: ")
			fmt.Scan(&u[index].Title)
			c++
		}

		if lsch == "content" || lsch == "both" {

			fmt.Println("\nEnter New Content: ")
			fmt.Scan(&u[index].Content)
			c++
		}

		// Ensuring valid field has been given
		if c > 0 {

			break
		} else {

			fmt.Println("\nPlease enter a valid field.......")
		}
	}

	// Updating updation time
	u[index].Updated_at = time.Now().Format("01-02-2006 15:04:05 Monday")

	// Adding the updated data to the []Header slice
	u = append(u, u[index])

	// Removing old data from the []Header slice
	u = append(u[:index], u[index+1:]...)

	// Feeding updated data to the filed
	feed(u)

	// Printing success message
	fmt.Println("\nNOTE Updated Successfully......")

	// Asking whether user wants to update another note
	var ch string

	for {

		fmt.Print("\nDo you wish to update another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		// Redirecting to update() through distro()
		case "y", "Y":
			distro(vomit(), "update")
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
