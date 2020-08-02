package main

import "log"

func main() {
	err := postTwitter()
	if err != nil {
		log.Fatal(err)
	}

	err = postRequest()
	if err != nil {
		log.Fatal(err)
	}
}
