# error

Go tem um sistema muito parecido com C de tratamento de erros.

```
f, err := os.Open("filename.ext")
if err != nil {
    panic(err)
}
```

---
[Inicio](README.md)
