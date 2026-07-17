package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type User struct {
	XMLName xml.Name `json:"-" xml:"user"`
	ID      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	Email   string   `json:"email,omitempty" xml:"email,omitempty"`
}

func DecodeJSON(data []byte) (User, error) {
	var user User
	err := json.Unmarshal(data, &user)
	return user, err
}

func EncodeJSON(user User) ([]byte, error) {
	return json.MarshalIndent(user, "", "  ")
}

func EncodeXML(user User) ([]byte, error) {
	return xml.MarshalIndent(user, "", "  ")
}

func main() {
	data := []byte(`{"id":7,"name":"Ada","email":"ada@example.com"}`)
	user, err := DecodeJSON(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode json: %v\n", err)
		os.Exit(1)
	}

	jsonData, err := EncodeJSON(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode json: %v\n", err)
		os.Exit(1)
	}

	xmlData, err := EncodeXML(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode xml: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
	fmt.Println(string(xmlData))
}
