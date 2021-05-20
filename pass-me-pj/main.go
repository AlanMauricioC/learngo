package main

import (
	"fmt"
	"os"
)

const (
	usage   = "Usage [user] [password]"
	accesOk = "Access granted to %q \n"
	errUser = "Access denied for %q \n"

	user, user2         = "inanc", "jack"
	password, password2 = "1879", "1888"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}

	usr := os.Args[1]
	psw := os.Args[2]

	if (user == usr && password == psw) || (user2 == usr && password2 == psw) {
		fmt.Printf(accesOk, usr)
		return
	}

	fmt.Printf(errUser, usr)

}
