package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//contact type data structure
type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Email string `json:"email"`
		Mobile string `json:"mobile"`
	} `json:"contact"`
}

//json data
var JSON = `{
	"name": "Babu",
	"title": "Shopkeeper",
	"contact": {
		"email": "babu@email.com",
		"mobile": "345.450.9900"
	}
}`

func main() {
	//var c Contact
	//Unmarshall into map variable JSON data
	var c map[string]interface{}
	 err := json.Unmarshal([]byte (JSON), &c)
	 if err != nil {
		log.Println("Error: ", err)
		return
	 }
	 fmt.Println(c)

}