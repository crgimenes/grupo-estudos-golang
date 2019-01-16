# Serial

Em exemplo simples de como ler a porta serial usando o pacote tarm/serial.

## instalar o pacote

```
go get github.com/tarm/serial
```

## Abre a porta serial

```go
c := &serial.Config{Name: "/dev/porta-serial", Baud: 115200}
s, err := serial.OpenPort(c)
if err != nil {
	log.Fatal(err)
}
```

## Lendo

```go
buf := make([]byte, 128)
n, err := s.Read(buf)
if err != nil {
	log.Fatal(err)
}
log.Print(string(buf[:n]))
```

## Fechando

É muito importante sempre fechar a porta serial porque esse se ela ficar aberta nenhum programa vai conseguir usar ela.

```go
err = s.Close()
if err != nil {
	log.Fatal(err)
}
```
