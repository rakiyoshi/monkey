package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(user)
	}
	fmt.Printf("Hello %s! This is the Monkey programing language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
