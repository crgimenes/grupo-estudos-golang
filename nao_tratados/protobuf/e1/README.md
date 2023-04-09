# Protocol Buffers

Instalando plugin go

```console
go get -u github.com/golang/protobuf/protoc-gen-go
````

Gerando codigo direto chamando `protoc`

```console
protoc --go_out=. ./user/user.proto
```

Gerando codigo usando `go generate`

```console
go generate
```

Para rodar o exemplo basta rodar `go run main.go`