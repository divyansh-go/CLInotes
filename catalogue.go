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

	for i, val := range u {

		fmt.Println(i+1, "  ", val.Id, "  ", val.Title)
	}

}
