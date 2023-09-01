# beanstalk

beanstalk é um servidor de filas bastante estável e rápido, e o melhor de tudo, com uma API fácil de conectar com suporte a Go.

instruções para instalação você encontra aqui: https://kr.github.io/beanstalkd/download.html

## aurora

Aurora é um visualizador web para você gerenciar suas filas no beanstalk, escrito em go e fácil de usar, veja como instalar aqui: https://github.com/xuri/aurora

## Exemplos

### Enviando dados para a fila

```go
package main

import (
    "fmt"
    "time"

    "github.com/kr/beanstalk"
)

func main() {
    c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        err = c.Close()
        if err != nil {
            fmt.Println()
        }
    }()

    id, err := c.Put([]byte("Olá mundo!"), 1, 0, time.Duration(120)*time.Second)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(id)
}
```

### Recebendo dados da fila

```go
package main

import (
    "fmt"
    "time"

    "github.com/kr/beanstalk"
)

func main() {
    c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        err = c.Close()
        if err != nil {
            fmt.Println()
        }
    }()

    id, body, err := c.Reserve(5 * time.Second)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(id, string(body))

    err = c.Delete(id)
    if err != nil {
        fmt.Println(err)
    }
}
```