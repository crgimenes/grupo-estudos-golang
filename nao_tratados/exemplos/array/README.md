# Arrays em Golang

Em Go arrais tem sempre tamanho fixo. 

Quando você adiciona um novo item em um array o que acontece internamente é que um novo array será criado contendo o array antigo e os novos itens. Eventualmente, o *garbage collector* vai liberar a memória alocada pelo array antigo.

Por isso, é importante tomar cuidado, é muito mais eficiente criar o array já com o tamanho necessário sempre que for possível e evitar realocações. Principalmente se for adicionar itens em um array usando um loop.

## Exemplos

Declarando um array de inteiros e atribuindo valores.

```golang
var a [3]int

a[0] = 1
a[1] = 2
a[2] = 3
```

Declarando um array usando forma de declaração curta e já atribuindo valores.

```golang
a2 := [3]int{4, 5, 6}
```

Você também pode atribuir os valores declarando com *var*.

```golang
var a3 = [3]int{7, 8, 9}
```

Também é possível declarar array especificando as posições dos elementos.

```golang
var a4 = [3]int{2: 12, 1: 11, 0: 10}
```

Não é necessário especificar o tamanho do array se você preencher os elementos, você pode usar o operador elipse, onde quem determina o tamanho do array é a quantidade de elementos informados.

```golang
a5 := [...]int{13, 14, 15}
```

Teste esses exemplos no [Go Playground](https://play.golang.org/p/YmmfIIFO_By)

