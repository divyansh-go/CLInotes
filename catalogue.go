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

func list(u []Header) {

	fmt.Print("\n")
	for i, val := range u {

		fmt.Println(i+1, "  ", val.Id, "  ", val.Title)
	}

}

func view(u []Header) {

	list(u)

	var ch int
	fmt.Print("\nEnter note ID: ")
	fmt.Scan(&ch)

	for _, val := range u {

		if val.Id == ch {
			fmt.Println("\nId: ", val.Id)
			fmt.Println("Title: ", val.Title)
			fmt.Println("\nContent:")
			fmt.Println(val.Content)
			fmt.Println("\nCreated_at: ", val.Created_at)
			fmt.Println("Updated_at: ", val.Updated_at)
		}
	}
}

func delete(u []Header) {

	var ch, index int
	fmt.Print("\nEnter note Id: ")
	fmt.Scan(&ch)

	fmt.Println("\nYou are about to delete:")
	for i, val := range u {

		if val.Id == ch {
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
		delete(vomit())
	}

}
