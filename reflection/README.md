# Reflection

Reflection permite obter informações sobre artefatos instanciados, é muito útil quando estamos trabalhando com interfaces.

Exemplo:

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name, omitempty"`
	Email    string `json:"email, omitempty"`
	Password string `json:"password, omitempty"`
}

func main() {

	// Pegando metadados da struct por reflection

	var u User

	f := reflect.Indirect(reflect.ValueOf(u))

	for i := 0; i < f.NumField(); i++ {

		fmt.Println("Name:", f.Type().Field(i).Name)
		fmt.Println("Tag:", f.Type().Field(i).Tag)
		fmt.Println("Index:", f.Type().Field(i).Index)
		fmt.Println("Offset:", f.Type().Field(i).Offset)
		fmt.Println("Type:", f.Type().Field(i).Type)
		fmt.Println("-=-=-=-=-=-=-=-")

	}

	// Pegando metadados da interface por reflection

	var iface interface{}

	iface = u

	f = reflect.Indirect(reflect.ValueOf(iface))

	for i := 0; i < f.NumField(); i++ {
		fmt.Println("Name:", f.Type().Field(i).Name)
		fmt.Println("Tag:", f.Type().Field(i).Tag)
		fmt.Println("Index:", f.Type().Field(i).Index)
		fmt.Println("Offset:", f.Type().Field(i).Offset)
		fmt.Println("Type:", f.Type().Field(i).Type)
		fmt.Println("-=-=-=-=-=-=-=-")
	}
}
```
[Playground](https://play.golang.org/p/0v0KLs3UuN)


---
[Inicio](../README.md)

[< interface](../interface/) - [error >](../error/)
