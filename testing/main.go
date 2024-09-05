package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Lire le contenu du fichier
	contenu, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	contenuString := string(contenu)
	fmt.Println(Split(contenuString))
}

func Split(str string) []string {
	slice := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "," {
			if word != "" {
				slice = append(slice, word)
				slice = append(slice, "-")
				word = ""
			}
			slice = append(slice, string(str[i]))
			slice = append(slice, "-")
			continue
		}
		if string(str[i]) != " " && string(str[i]) != "\t" {
			word += string(str[i])
		} else if word != "" {
			slice = append(slice, word)
			slice = append(slice, "-")
			word = ""
		}
	}
	slice = append(slice, word)

	return slice
}
