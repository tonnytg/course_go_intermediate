package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Request")

	resp, err := http.Get("https://646a8f077d3c1cae4ce2a7fc.mockapi.io/api/v1/users")
	if err != nil {
		fmt.Println("Aconteceu algum erro:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body)
}
