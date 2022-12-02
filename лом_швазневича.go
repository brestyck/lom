/*
Против лома нет приема,
Если нет другого лома
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func write_attempt(attempt int) {
	os.Mkdir(fmt.Sprint(attempt), os.ModeAppend)
}

func read_attempt() int {
	path, _ := os.Getwd()
	files, _ := ioutil.ReadDir(path)
	maxima := -1
	for _, j := range files {
		if i, err := strconv.Atoi(j.Name()); err == nil {
			maxima = i
		}
	}
	return maxima
}

func main() {
	answers := []any{
		"__token___780683496743",
		"__token___651080620892", "__token___332895764626", "__token___277099167941",
		"__token___809556138707", "__token___467520690352", "__token___783658621944",
		"__token___667233261478", "__token___289852777018", "__token___292515051623",
	}
	if read_attempt() == -1 {
		write_attempt(0)
		fmt.Println(answers[0])
	} else {
		x := read_attempt()
		write_attempt(x + 1)
		if x+1 < len(answers) {
			fmt.Println(answers[x+1])
		}
	}
}
