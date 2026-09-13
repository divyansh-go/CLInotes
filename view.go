package main

import "fmt"

func view(u []Header, ch1 int) {

	// Printing given Header var in proper format
	for _, val := range u {

		if val.Id == ch1 {
			fmt.Println("\n----------------------------------------------")
			fmt.Println("\nId: ", val.Id)
			fmt.Println("Title: ", val.Title)
			fmt.Println("\nContent:")
			fmt.Println(val.Content)
			fmt.Println("\nCreated_at: ", val.Created_at)
			fmt.Println("Updated_at: ", val.Updated_at)
			fmt.Println("\n----------------------------------------------")
		}
	}

}

// Asks user if they wish to view another Note, after executing view()
func recurrview() {

	var ch string

	for {

		fmt.Print("\nDo you wish to view another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(vomit(), "view")
			return

		case "n", "N":
			distro(vomit(), menu())
			return

		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}
}
