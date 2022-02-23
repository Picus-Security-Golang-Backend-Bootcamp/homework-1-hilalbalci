package main

import (
	"fmt"
	"hw1/m/myfunctions"
	"os"
	"strings"
)

func main() {
	checkArgs()
}

func checkArgs() {
	myList := []string{"In Search of Lost Time",
		"Ulysses",
		"Ul",
		"Don Quixote",
		"Don Quixote Second",
		"The Great Gatsby",
		"Moby Dick",
		"Hamlet",
		"The Odyssey",
		"Madame Bovary"}
	if (len(os.Args[1:]) == 0) || ((strings.ToLower(os.Args[1]) != "search") && (strings.ToLower(os.Args[1]) != "list")) {
		fmt.Print("Bu uygulamada kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n")
	} else {
		myfunctions.SearchOrList(strings.Join(os.Args[2:], " "), strings.ToLower(os.Args[1]) == "search", myList)
	}
}
