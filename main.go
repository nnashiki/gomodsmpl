package main

import (
	"fmt"
	"os"
	"strconv"
	"gomodsmpl/users"
)

func main() {
	var age int
	var age_arg string
	var err interface{}
	age_arg = os.Args[3]
	name := os.Args[2]

	age, err = strconv.Atoi(age_arg)
	if err != nil {
		fmt.Print("error が出ました")
	}

	var user users.User
	if os.Args[1] == "post" {
		user, err = users.Make_user(name, age)
	}
	fmt.Print(user.Name)
	fmt.Print(user.Age)
}

