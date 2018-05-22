package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {
	slovo := Scan1()

	type City struct {
		Peter, Moscow
	}

	if slovo == 

}
