package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	distro(menu())

}

func menu() string {

	var choice string

	fmt.Println("\n Options Available:\n ~ Add\n ~ List\n ~ View\n ~ Update\n ~ Delete\n ~ Help\n ~ Exit")
	fmt.Print("\n\nChoose an action to perform: ")
	fmt.Scan(&choice)

	return choice
}

func distro(r1 string) {

	lchoice := strings.ToLower(r1)

	switch lchoice {
	case "add":
		add()
	case "list":
		list(vomit())
	case "view":
		view(vomit())
	case "update":
		fmt.Println("responding")
	case "delete":
		fmt.Println("responding")
	case "help":
		fmt.Println("responding")
	case "exit":
		fmt.Println("responding")
	default:
		fmt.Println("\nInvalid Action. Please choose a valid action to perform......")
		distro(menu())
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
