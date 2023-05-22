package web

import (
	"consumer/internal/domain"
	"encoding/json"
	"io"
	"net/http"
)

func GetUsers() (domain.Users, error) {
	resp, err := http.Get("https://646a8f077d3c1cae4ce2a7fc.mockapi.io/api/v1/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users domain.Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
