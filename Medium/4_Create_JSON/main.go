package main

import (
	"encoding/json"
	"fmt"
)

// `json: "<key>"` will show the property name as the key value
type course struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"password"`
	Tags []string `json:"tags"`
}

func EncodeJson(){
	courseData := []course{
		{"ReactJS", 299, "learncode.com", "react123", []string{"web-dev", "js"}},
		{"angularJS", 399, "learncode.com", "angular123", []string{"web-dev", "js", "frontend"}},
		{"NodeJS", 499, "learncode.com", "node123", []string{"backend", "js"}},
	}

	jsonData, err := json.MarshalIndent(courseData, "", "\t")
	// MarshalIndent indents the json data and Marshal does not indent the json data

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}

func main(){
	EncodeJson()
}