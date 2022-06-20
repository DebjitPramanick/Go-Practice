package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl string = "http://localhost:8000/create"

func main(){

	requestBody := strings.NewReader(`
		{
			"courseName": "Go Language",
			"price": 1000
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Method 1 to get data
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}