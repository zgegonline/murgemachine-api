package main

import (
	"fmt"
	"log"

	"github.com/zgegonline/murgemachine-restapi/api"
)

func main() {
	fmt.Println("Hello World !")

	log.Fatal(api.Start())
}
