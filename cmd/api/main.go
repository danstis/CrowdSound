package main

import (
	"log"

	"github.com/danstis/CrowdSound/pkg/api"
)

const (
	ADDRESS string = ":9200"
)

func main() {
	api := api.New()
	api.Address = ADDRESS
	log.Fatal(api.Run())
}
