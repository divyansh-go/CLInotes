package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func add() {

	// Taking Inputs
	var tit, cont string

	randId := rand.Intn(9000) + 1000

	fmt.Print("\nTitle : ")
	fmt.Scan(&tit)

	fmt.Print("\nContent : ")
	fmt.Scan(&cont)

	ctime := time.Now().Format("01-02-2006 15:04:05 Monday")

	fmt.Println("\n\nNOTE Created Successfully at ", ctime)
	fmt.Println("Id ~ ", randId)
	fmt.Println("Title ~ ", tit)

	user := Header{
		Id:         randId,
		Title:      tit,
		Content:    cont,
		Created_at: ctime,
		Updated_at: ctime,
	}

	// Writing Data
	feed(user)

	// Recall
	var ch string

	for {

		fmt.Print("\nDo you wish to add another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(nil, "add")
			return

		case "n", "N":
			distro(vomit(), menu())
			return

		default:
			fmt.Println("\nInvalid Input. Please try again.....")
		}
	}
}

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

func view(u []Header, ch1 int) {

	// Printing struct in proper format
	for _, val := range u {

		if val.Id == ch1 {
			fmt.Println("\nId: ", val.Id)
			fmt.Println("Title: ", val.Title)
			fmt.Println("\nContent:")
			fmt.Println(val.Content)
			fmt.Println("\nCreated_at: ", val.Created_at)
			fmt.Println("Updated_at: ", val.Updated_at)
		}
	}

}

func delete(u []Header, ch1 int) {

	// Printing info of the chosen noteid
	// Taking index value of the chosen noteid
	var index int

	fmt.Println("\nYou are about to delete:")
	for i, val := range u {

		if val.Id == ch1 {
			fmt.Println("\nId: ", val.Id)
			fmt.Println("Title: ", val.Title)
			index = i
		}
	}

	var sch string
	fmt.Print("\nAre you sure? (y/n): ")
	fmt.Scan(&sch)
	lsch := strings.ToLower(sch)

	switch lsch {
	case "y":
		u = append(u[:index], u[index+1:]...)

		jdata, _ := json.Marshal(u)
		os.WriteFile("notes.json", jdata, 0666)
	case "n":
		delete(u, ask())
	}

	var ch string

	for {

		fmt.Print("\nDo you wish to delete another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(nil, "delete")
			return

		case "n", "N":
			distro(vomit(), menu())
			return

		default:
			fmt.Println("\nInvalid Input. Please try again.....")
		}
	}
}

func update(u []Header, ch1 int) {

	view(u, ch1)

	var sch string
	fmt.Print("\nEnter the field to be updated. (Title/Content/Both): ")
	fmt.Scan(&sch)
	lsch := strings.ToLower(sch)

	var index int

	for i, val := range u {

		if val.Id == ch1 {

			index = i
		}
	}

	if lsch == "title" || lsch == "both" {
		fmt.Print("\nEnter New Title: ")
		fmt.Scan(&u[index].Title)
	}
	if lsch == "content" || lsch == "both" {
		fmt.Println("\nEnter New Content: ")
		fmt.Scan(&u[index].Content)
	}
	u[index].Updated_at = time.Now().Format("01-02-2006 15:04:05 Monday")

	u = append(u, u[index])

	u = append(u[:index], u[index+1:]...)

	jdata, _ := json.Marshal(u)
	os.WriteFile("notes.json", jdata, 0666)

	var ch string

	for {

		fmt.Print("\nDo you wish to update another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(nil, "update")
			return

		case "n", "N":
			distro(vomit(), menu())
			return

		default:
			fmt.Println("\nInvalid Input. Please try again.....")
		}
	}
}

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

func recurrview() {

	var ch string

	for {

		fmt.Print("\nDo you wish to view another note? (y/n): ")
		fmt.Scan(&ch)

		switch ch {

		case "y", "Y":
			distro(nil, "view")
			return

		case "n", "N":
			distro(vomit(), menu())
			return

		default:
			fmt.Println("\nInvalid Input. Please try again.....")
		}
	}
}
