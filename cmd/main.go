package main

import "fmt"

func main() {
	_ = GetText("batata")
	_ = GetText("macarrão")
	_ = GetText("arroz")
	_ = GetText("feijão")
}

func GetText(a string) string {
	return a
}
