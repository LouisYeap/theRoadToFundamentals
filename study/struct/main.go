package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	//里面的Person可以省略
	persons := []Person{{Age: 10, Sex: "男", Name: "John"}, {Age: 20, Sex: "男", Name: "mike"}}
	//fmt.Println(persons)

	fmt.Printf("%+v\n", persons)

	a := struct {
		Name string
	}{
		Name: "John",
	}
}
