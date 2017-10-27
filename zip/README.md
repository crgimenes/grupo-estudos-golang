# Zip

Como comprimir e descomprimir arquivos zip


## Comprimir 

```go
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    // arquivos para comprimir
    files := []string{
        "teste1.txt",
        "teste2.txt",
        "teste3.txt",
    }

    // cria o arquivo .zip
    zipFile, err := os.Create("arquivo.zip")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer zipFile.Close() // quando terminar fecha o arquivo .zip

    // cria um writer que vai escrever no arquivo .zip
    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close() // quando terminar fecha o writer

    for i, file := range files {

        fmt.Printf("comprimindo arquivo #%02d %v\n", i+1, file)

        // abre o arquivo que vai ser comprimido
        var fileToCompress *os.File
        fileToCompress, err = os.Open(file)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer fileToCompress.Close()

        // pega infos do arquivo que vai ser comprimido
        var info os.FileInfo
        info, err = fileToCompress.Stat()
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("tamanho original %d bytes\n", info.Size())

        // prepara informações do arquivo que vai ser comprimido
        // para colocar no cabeçalho do arquivo .zip
        var header *zip.FileHeader
        header, err = zip.FileInfoHeader(info)
        if err != nil {
            fmt.Println(err)
            return
        }

        // ajusta metodo de compressão
        //header.Method = zip.Store
        header.Method = zip.Deflate

        // grava cabeçalho do zip
        var writer io.Writer
        writer, err = zipWriter.CreateHeader(header)
        if err != nil {
            fmt.Println(err)
            return
        }

        // grava arquivo comprimido no arquivo .zip
        _, err = io.Copy(writer, fileToCompress)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}
```

## Descomprimir 

```go
package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {
    // abre arquivo .zip
    zipFile, err := zip.OpenReader("arquivo.zip")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer zipFile.Close()

    for i, file := range zipFile.File {

        fmt.Printf("descomprimindo arquivo #%02d %v\n", i+1, file.Name)

        // abre reader para ler arquivo de dentro do zip
        reader, err := file.Open()
        if err != nil {
            fmt.Println(err)
            return
        }
        defer reader.Close()

        var f *os.File
        // abre arquivo de destino
        f, err = os.OpenFile(
            file.Name,
            os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
            file.Mode())
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()

        // grava arquivo de destino
        _, err = io.Copy(f, reader)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}
```

---
[Inicio](../README.md)

[< crc](../crc/) - [wikipedia >](./wikipedia/)
