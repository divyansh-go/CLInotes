package main

import (
	"encoding/json"
	"os"
)

func feed(slc []Header) {

	jdata, _ := json.Marshal(slc)
	os.WriteFile("notes.json", jdata, 0666)

}

func vomit() []Header {

	var users []Header

	jdata, _ := os.ReadFile("notes.json")
	json.Unmarshal(jdata, &users)

	return users
}
