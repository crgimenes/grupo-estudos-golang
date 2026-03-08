package main

import (
	"fmt"
	"os"
)

func main() {
	content := "package main\n\nconst generatedMessage = \"generated file ready\"\n"
	if err := os.WriteFile("generated_message.go", []byte(content), 0o644); err != nil {
		panic(err)
	}
	fmt.Println("generated generated_message.go")
}
