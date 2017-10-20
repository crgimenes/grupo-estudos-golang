# CRC

Calcular o CRC de uma string ou um arquivo é muito útil para garantir que algum dado não foi alterado. As funções CRC tem a vantagem de serem rápidas comparado com outras formas de hash.

```
package main

import (
    "fmt"
    "hash/crc32"
)

func main() {
    valor := "Isto é um teste"    
    checksum := crc32.ChecksumIEEE([]byte(valor))
    fmt.Printf("Checksum: 0x%x\n", checksum)
}
```

---
[Inicio](../README.md)

[< data-race](../data-race/) - [zop >](../zip/)
