package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl string = "http://localhost:8000/read"

func main(){
	fmt.Println("Handling URL.")
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	// Method 1 to get data
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// Method 2 to get data
	var responseString strings.Builder
	responseString.Write(content)
	fmt.Println(responseString.String())
}