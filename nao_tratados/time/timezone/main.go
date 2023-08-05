package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		s string
		t time.Time
	)

	t = time.Now() // Data e hora atual

	// Data e hora atual UTC
	s = t.UTC().Format("2006-01-02 15:04:05") // formato da data

	fmt.Printf("Data e hora atual UTC............: %s\n", s)

	// Data e hora atual local (timezone do sistema)
	s = t.Format("2006-01-02 15:04:05") // formato da data

	fmt.Printf("Data e hora atual local..........: %s\n", s)

	// Data com timezone de Sao Paulo
	loc, _ := time.LoadLocation("America/Sao_Paulo") // prepara o timezone
	s = t.In(loc).Format("2006-01-02 15:04:05")      // formato da data

	fmt.Printf("Data e hora atual São Paulo......: %s\n", s)

}
