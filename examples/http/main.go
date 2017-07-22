package main

import (
	"log"
	"net/http"

	"github.com/mirza-s/uarand"
)

func main() {
	req, err := http.NewRequest("POST", "https://requestb.in/t0wbsbt0", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", uarand.GetRandom())

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
