package main

import (
	"log"

	"github.com/viniciusbls9/go-movie/pkg/utils"
)

func main() {
	err := utils.CreateRouters()
	if err != nil {
		log.Fatal(err)
	}
}
