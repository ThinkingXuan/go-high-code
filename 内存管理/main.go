package main

import "fmt"

type user struct {
	usename  string
	password string
}

func main() {
	u := user{}
	s := u.usename
	b := u.password
	fmt.Println(s, b)
}
