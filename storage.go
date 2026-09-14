package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func feed(slc []Header) {

	jdata, err := json.Marshal(slc)

	if err != nil {

		fmt.Println("\nFailed to convert NOTE to JSON.")
		fmt.Println("Error: ", err)

		return
	}

	err = os.WriteFile("notes.json", jdata, 0666)

	if err != nil {

		fmt.Println("\nFailed to write file.")
		fmt.Println("Error: ", err)

		return
	}

}

func vomit() []Header {

	var users []Header

	jdata, err := os.ReadFile("notes.json")

	if err != nil {

		fmt.Println("\nFailed to read file.")
		fmt.Println("Error: ", err)

		return nil
	}

	err = json.Unmarshal(jdata, &users)

	if err != nil {

		fmt.Println("\nFailed to convert NOTE to JSON.")
		fmt.Println("Error: ", err)

		return nil
	}

	return users
}
