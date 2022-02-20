package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
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
