package main

import (
	"fmt"
	"math/rand"
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
