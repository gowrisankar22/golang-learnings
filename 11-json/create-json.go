package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Perms map[string]bool

type user struct {
	FirstName string `json:"first_name"`
	Password  string `json:"-"`
	Perms     `json:"perms,omitempty"`
}

func main() {

	users := []user{
		{
			FirstName: "Ajay",
			Password:  "abc1234",
			Perms:     Perms{"admin": true},
		},
		{
			FirstName: "Rahul",
			Password:  "xyz",
		},
		{
			FirstName: "John",
			Password:  "abc1234",
			Perms:     Perms{"write": false},
		},
	}

	//jsonData, err := json.Marshal(users)
	jsonData, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Fatal("Program stopped with an", err)
	}

	err = ioutil.WriteFile("user.json", jsonData, 666)

	if err != nil {
		log.Fatal("Program stopped with an", err)
	}

}
