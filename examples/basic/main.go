package main

import (
	"log"

	"github.com/mirza-s/uarand"
)

func main() {
	log.Printf("%s\n", uarand.GetRandom())
	log.Printf("%s\n", uarand.GetRandom())
	log.Printf("%s\n", uarand.GetRandom())
}
