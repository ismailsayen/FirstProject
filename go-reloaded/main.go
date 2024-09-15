package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	goReloaded "goReloaded/functions"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var contenuString []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contenuString = append(contenuString, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	sl := [][]string{}
	for i := 0; i < len(contenuString); i++ {
		sl = append(sl, goReloaded.Split(contenuString[i]))
	}
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			wd := string(sl[i][j])
			if wd == "(bin)" && j != 0 {
				a, _ := strconv.ParseInt(string(sl[i][j-1]), 2, 64)
				sl[i][j-1] = strconv.Itoa(int(a))
				sl[i] = append(sl[i][:j], sl[i][j+1:]...)
			} else if wd == "(hex)" && j != 0 {
				a, _ := strconv.ParseInt(string(sl[i][j-1]), 16, 64)
				sl[i][j-1] = strconv.Itoa(int(a))
				sl[i] = append(sl[i][:j], sl[i][j+1:]...)
				j = 0
			} else if wd == "(up)" && j != 0 {
				sl[i][j-1] = goReloaded.Upper(sl[i][j-1])
				sl[i] = append(sl[i][:j], sl[i][j+1:]...)
				j = 0
			} else if wd == "(low)" && j != 0 {
				sl[i][j-1] = goReloaded.LowerCase(sl[i][j-1])
				sl[i] = append(sl[i][:j], sl[i][j+1:]...)
				j = 0
			} else if wd == "(cap)" && j != 0 {
				sl[i][j-1] = goReloaded.Capitalize(sl[i][j-1])
				sl[i] = append(sl[i][:j], sl[i][j+1:]...)
				j = 0
			}
			if wd == "(up," && (j != 0 && j != len(sl[i])-1) {
				ve := sl[i][j+1]
				nb := ve[0 : len(ve)-1]
				if ve[len(ve)-1] == ')' && goReloaded.Is_number(nb) {
					n, _ := strconv.Atoi(nb)
					if n >= j {
						n = j
					}
					for k := j - n; k < j; k++ {
						sl[i][k] = goReloaded.Upper(sl[i][k])
					}
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					j = 0
				}
			} else if wd == "(low," && (j != 0 && j != len(sl[i])-1) {
				ve := sl[i][j+1]
				nb := ve[0 : len(ve)-1]
				if ve[len(ve)-1] == ')' && goReloaded.Is_number(nb) {
					n, _ := strconv.Atoi(nb)
					if n >= j {
						n = j
					}
					for k := j - n; k < j; k++ {
						sl[i][k] = goReloaded.LowerCase(sl[i][k])
					}
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					j = 0
				}
			} else if wd == "(cap," && (j != 0 && j != len(sl[i])-1) {
				ve := sl[i][j+1]
				nb := ve[0 : len(ve)-1]
				if ve[len(ve)-1] == ')' && goReloaded.Is_number(nb) {
					n, _ := strconv.Atoi(nb)
					if n >= j {
						n = j
					}
					for k := j - n; k < j; k++ {
						sl[i][k] = goReloaded.Capitalize(sl[i][k])
					}
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					sl[i] = append(sl[i][:j], sl[i][j+1:]...)
					j = 0
				}
			}
		}
	}
	for i := 0; i < len(sl); i++ {
		sl[i] = goReloaded.SplitWithPunc(sl[i])
	}

	result := ""
	for i := 0; i < len(sl); i++ {
		result += goReloaded.Printing(sl[i])
	}
	if len(result) > 0 {
		result = result[0 : len(result)-1]
	}
	fmt.Println(result)
}
