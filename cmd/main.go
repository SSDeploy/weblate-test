package main

import "fmt"

func main() {
	fmt.Println(GetText("batata"))
	fmt.Println(GetText("macarrão"))
	fmt.Println(GetText("arroz"))
	fmt.Println(GetText("feijão"))
}

func GetText(a string) string {
	return a
}
