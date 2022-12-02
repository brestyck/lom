/*
Против лома нет приема,
Если нет другого лома
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func write_attempt(attempt int) {
	os.Mkdir(fmt.Sprint(attempt), os.ModeAppend)
}

func read_attempt() int {
	path, _ := os.Getwd()
	files, _ := os.ReadDir(path)
	maxima := -1
	for _, j := range files {
		if i, err := strconv.Atoi(j.Name()); err == nil {
			maxima = i
		}
	}
	return maxima
}

func main() {
	answers := []any{}
	if read_attempt() == -1 {
		write_attempt(0)
		fmt.Println(answers[0])
	} else {
		x := read_attempt()
		if x > len(answers)-1 {
			fmt.Println(answers[0])
		} else {
			write_attempt(x + 1)
			fmt.Println(answers[x+1])
		}
	}
}
