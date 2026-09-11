package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	distro(menu())

}

func menu() string {

	var choice string

	fmt.Println("\n Options Available:\n ~ Create\n ~ List\n ~ View\n ~ Update\n ~ Delete\n ~ Help\n ~ Exit")
	fmt.Print("\n\nChoose an action to perform: ")
	fmt.Scan(&choice)

	return choice
}

func distro(r1 string) {

	lchoice := strings.ToLower(r1)

	switch lchoice {
	case "create":
		create()
	case "list":
		fmt.Println("responding")
	case "view":
		fmt.Println("responding")
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

	jdata, _ := json.Marshal(struc)

	file, _ := os.OpenFile("notes.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	file.Write(jdata)

	defer file.Close()

}

func create() {

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

	feed(user)
}
