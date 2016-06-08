# defer

A clausula defer define que uma função deve ser chamada no fim da execução da rotina atual.
Essa clausula é muito útil para por exemplo fechar arquivos abertos durante a execução da função.

```
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

```

---
[Inicio](README.md)
