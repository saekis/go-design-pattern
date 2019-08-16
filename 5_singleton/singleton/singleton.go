package singleton

import "fmt"

type singleton struct {
}

var instance = newInstance()

func newInstance() *singleton {
	fmt.Println("Create new instance")
	return &singleton{}
}

func GetInstance() *singleton {
	fmt.Println("Exist instance")
	return instance
}
