package main

import "fmt"

func PrintWithCoolFormat(i interface{}, detail string) {
	fmt.Println("")
	fmt.Println("========================")
	fmt.Println(detail)
	fmt.Println("------------------------")
	fmt.Println(i)
	fmt.Println("------------------------")
	fmt.Println("========================")
	fmt.Println("")
}
