package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name	string	`json:"name"`
	Title	string	`json:"title"`
	Contact	struct {
		Home	string `json:"home"`
		Cell	string `json:"cell"`
	} `json:"contact"`
}

// JSON contains a sample string to unmarshal
var JSON = `{
    "name": "Gopher",
    "title": "programmer",
    "contact": {
        "home": "415.333.3333",
        "cell": "415.555.5555"
     }
}`

func main() {
	// Unmarshall the JSON string into our variable
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)
	fmt.Println(c.Name)
	fmt.Println(c.Title)
	fmt.Println(c.Contact.Home)
	fmt.Println(c.Contact.Cell)
}
