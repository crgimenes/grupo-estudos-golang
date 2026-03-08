package main

import (
	"encoding/json"
	"encoding/xml"
)

type User struct {
	ID    int     `json:"id" xml:"id"`
	Name  string  `json:"name" xml:"name"`
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func DecodeJSON(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}

func EncodeXML(u User) ([]byte, error) {
	return xml.Marshal(u)
}
