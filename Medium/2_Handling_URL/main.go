package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=golang&paymentId=P20062022"

func main(){
	fmt.Println("Handling URL.")
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Query())
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)

	// Get query params
	query_params := result.Query()
	fmt.Println("Payment ID ==> ", query_params.Get("paymentId"))

	// Create URL
	newURLParts := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/sample",
		RawPath: "user=debjit",
	}

	newURL := newURLParts.String()
	fmt.Println(newURL)
}