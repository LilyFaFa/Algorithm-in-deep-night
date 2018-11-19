package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Title string
}

func (this User) Test(x int) {
	fmt.Println("User.Test:", this.Name, x)
}

func test(i interface{}) {
	t := reflect.TypeOf(i)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println(k)
		return
	}
	// 输出 Type: User Kind: struct
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Type:", t.Name())
	// 输出
	/*
		Id:1
		Name:Lily
		Age:12
		Title:test
	*/
	v := reflect.ValueOf(i)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s:%v\n", field.Name, value)
	}

	/*
		main.User
		{1 Lily 12 test}
		{0  0 }
	*/
	fmt.Println(t)
	fmt.Println(v)
	cpy := reflect.New(v.Type()).Elem()
	fmt.Println(cpy)
}

func main() {
	u := User{1, "Lily", 12, "test"}
	test(u)
}
