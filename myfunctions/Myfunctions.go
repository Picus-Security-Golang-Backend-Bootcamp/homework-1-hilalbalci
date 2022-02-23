package myfunctions

import (
	"fmt"
	"strings"
)

func SearchOrList(keyword string, search bool, list []string) {
	if search {
		fmt.Print(contains(list, keyword))
	} else {
		fmt.Print(strings.Join(list[:], "\n"))
	}
}

func contains(s []string, str string) string {
	result := []string{}
	for _, v := range s {
		if (strings.ToLower(v) == strings.ToLower(str)) || strings.Contains(strings.ToLower(v), strings.ToLower(str)) {
			result = append(result, v)
		}
	}
	return strings.Join(result, "\n")
}
