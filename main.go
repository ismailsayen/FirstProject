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
	fmt.Println(len(Split(contenuString)))
}

func Split(str string) []string {
	slice := []string{}
	word := ""
	cha := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && string(str[i]) != "\t" && !isPunctuation(string(str[i])) {
			if cha != "" {
				slice = append(slice, cha)
				slice = append(slice, "-")
				cha = ""
			}
			word += string(str[i])
		} else {
			if word != "" {
				slice = append(slice, word)
				slice = append(slice, "-")
				word = ""
			}
			if isPunctuation(string(str[i])) {
				cha += string(str[i])
			} else {
				if cha != "" {
					slice = append(slice, cha)
					slice = append(slice, "-")
					cha = ""
				}
			}
		}
	}
	if word != "" {
		slice = append(slice, word)
	}
	if cha != "" {
		slice = append(slice, cha)
	}
	return slice
}

func isPunctuation(s string) bool {
	if s == "," || s == "?" || s == "!" || s == "." || s == ":" || s == ";" {
		return true
	}
	return false
}
