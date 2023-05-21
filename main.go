package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Users []struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Age       string    `json:"Age"`
	ID        string    `json:"id"`
}

func main() {
	fmt.Println("Request")

	resp, err := http.Get("https://646a8f077d3c1cae4ce2a7fc.mockapi.io/api/v1/users")
	if err != nil {
		fmt.Println("Aconteceu algum erro:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Body: %s\n", body)

	var users Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Deu erro ao converter")
	}
	fmt.Println(users[0].Name)
	fmt.Println(users[1].Name)
	fmt.Println(users[2].Name)
}
