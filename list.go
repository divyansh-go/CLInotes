package main

import "fmt"

func list(u []Header) {

	fmt.Println("\n=============================")
	fmt.Println("   ID   |  TITLE")
	fmt.Println("=============================")
	fmt.Println("")

	// Printing appropriate fields
	for _, val := range u {

		fmt.Println(" ", val.Id, " | ", val.Title)
	}
}

// Asks user if they wish to continue to menu, after executing list()
func recurrlist() {

	var ch string

	for {

		fmt.Print("\nDo you wish to perform another action? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(vomit(), menu())
			return

		case "n", "N":
			for {

				fmt.Print("\nAre you sure you want to exit? (y/n): ")
				fmt.Scan(&ch)

				switch ch {

				case "y", "Y":
					distro(nil, "exit")
					return

				case "n", "N":
					distro(vomit(), menu())
					return

				default:
					fmt.Println("\nInvalid Input. Please try again.....")

				}
			}

		default:
			fmt.Println("\nInvalid Input. Please try again.....")

		}
	}
}
