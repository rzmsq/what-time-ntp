package main

import (
	"log"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
