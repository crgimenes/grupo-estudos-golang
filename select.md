# select


Select trabalha junto com canais para esperar retornos de canais específicos.

```go
select {
case msg1 := <-c1:
    fmt.Println("canal 1 retornou :", msg1)
case msg2 := <-c2:
    fmt.Println("canal 2 retornou :", msg2)
}
```

---
[Inicio](README.md)
