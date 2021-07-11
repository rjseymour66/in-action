package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	// Create a map of k/v pairs
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{} {
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// Marshall the map into JSON string
	data, err := json.MarshalIndent(c, "", "    ")
//	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}

