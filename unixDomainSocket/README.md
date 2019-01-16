# Unix domain socket

Unix Domain Sockets ou [IPC](https://en.wikipedia.org/wiki/Inter-process_communication) socket é uma forma muito pratica e segura de trocar informações entre processos. Essa forma de IPC usa um arquivo como endereço/name space no lugar de um IP e uma porta como seria em uma comunicação via rede.

Uma coisa importante para ter em mente é que como vamos usar um arquivo o servidor é responsável por ele, se não existir ele sera criado automaticamente mas se não existir você vai receber um erro com algo como "bind: address already in use" que significa que o arquivo já existe e o servidor não tem como reaproveitar um arquivo que já existe, o correto é fazer shutdown elegantemente e fechar e apagar o arquivo antes de derrubar o servidor. E dependendo do sistema pode ser interessante verificar se o arquivo já existe e apagar antes de subir o servidor. 

Apesar da facilidade, como usamos um arquivo como endereço não da para usar para trocar informação entre maquinas diferentes, e quem fica responsável por manter essa comunicação é o kernel, o arquivo é apenas um name space, nenhum byte vai ser mesmo escrito no arquivo ele vai ocupar zero espaço de disco, toda a comunicação acontece na RAM e gerenciada pelo kernel.

Outra coisa muito importante é que usar Unix domain socket é um recurso padrão de qualquer ambiente [POSIX](https://en.wikipedia.org/wiki/POSIX) mas não esta presente por padrão no Windows.

Para testes podemos também usar o [netcat](https://en.wikipedia.org/wiki/Netcat) 

## Servidor usando netcat

```console
nc -lU /tmp/echo.sock && rm /tmp/echo.sock
```

## Cliente usando netcat

```console
nc -U /tmp/echo.sock
```

## Exemplos

### Executando o servidor echo

```console
go run echo/server/main.go
```

Em outro console execute

```console
go run echo/client/main.go
```
