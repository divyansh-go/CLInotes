package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		add()

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

type Header struct {
	Id         int
	Title      string
	Content    string
	Created_at string
	Updated_at string
}

func feed(struc Header) {

	s := vomit()
	s1 := []Header{struc}
	s = append(s, s1...)

	jdata, _ := json.Marshal(s)
	os.WriteFile("notes.json", jdata, 0666)

}

func vomit() []Header {

	var users []Header

	jdata, _ := os.ReadFile("notes.json")
	json.Unmarshal(jdata, &users)

	return users
}

func ask() int {

	var ch int
	fmt.Print("\nEnter note ID: ")
	fmt.Scan(&ch)

	return ch
}
