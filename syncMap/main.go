package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	m := sync.Map{}

	go func(m *sync.Map) {
		for {
			m.Range(func(key interface{}, value interface{}) bool {
				fmt.Printf("goroutine key: %v -> valor: %v\n", key, value)
				<-time.After(1 * time.Millisecond)
				return true
			})
		}
	}(&m)

	// da um tempinho para carregar a go routine
	<-time.After(5 * time.Millisecond)

	m.Store("item 1", "valor item 1")
	m.Store("item 2", "valor item 2")
	m.Store("item 3", "valor item 3")
	m.Store("item 4", "valor item 4")
	m.Store("item 5", "valor item 5")
	m.Store("item 6", "valor item 6")

	valor, ok := m.Load("item 2")
	if ok {
		fmt.Println("item encontrado", valor.(string))
		<-time.After(5 * time.Millisecond)
	}

	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> valor: %v\n", key, value)
		return true
	})

	m.Delete("item 5")

	fmt.Println("\nItem 5 foi removido, confira se é verdade:")
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> valor: %v\n", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})

	novoItem, ok := m.LoadOrStore("item 7", "valor item 7")
	if ok {
		fmt.Println("item 7 já estava no mapa")
	}

	fmt.Printf("\nitem 7: %v\n", novoItem)

	fmt.Println("\nAgora tem item 7, confira se é verdade:")
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> valor: %v\n", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})
}
