package main

import (
	"fmt"
	"log"
	"os"
	"goReloaded"
)

func main() {
	contenu, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	contenuString := string(contenu)
	fmt.Println(goReloaded.Split(contenuString))
	fmt.Println(len(goReloaded.Split(contenuString)))
}

