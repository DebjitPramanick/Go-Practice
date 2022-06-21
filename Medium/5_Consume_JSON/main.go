package main

import (
	"encoding/json"
	"fmt"
)

// `json: "<key>"` will show the property name as the key value
type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags"`
}

func DecodeJson() {
	jsonData := []byte(`
        {
                "name": "ReactJS",
                "price": 299,
                "platform": "learncode.com",
                "password": "react123",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var newCourse course

	isValid := json.Valid(jsonData)

	if isValid {
		fmt.Printf("JSON is valid.\n")
		json.Unmarshal(jsonData, &newCourse)
		fmt.Printf("%#v\n", newCourse)
	} else {
		fmt.Println("JSON is not valid.")
	}
}

func main() {
	DecodeJson()
}
