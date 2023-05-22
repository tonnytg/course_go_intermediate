package main

import (
	"consumer/pkg/web"
	"fmt"
)

func main() {
	fmt.Println("Request")

	users, err := web.GetUsers()
	if err != nil {
		fmt.Println("Deu erro:", err)
	}

	for i, v := range users {
		fmt.Printf("[%d] User: %s\n", i, v.Name)
	}
}
