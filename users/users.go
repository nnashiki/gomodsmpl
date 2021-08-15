package users

import "errors"

type User struct {
	Name string
	Age  int
}

func Make_user(name string, age int) (User, error) {
	if name == "root" {
		return User{Name: "", Age: 0}, errors.New("hoge")
	}
	return User{Name: name, Age: age}, nil
}
