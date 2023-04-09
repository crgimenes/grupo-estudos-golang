# Operadores binários em Golang

Exemplos de operadores binários em Go

É importante não confundir operadores binários com operadores lógicos, operadores lógicos tem os mesmos nomes mas usam símbolos diferentes e geralmente agem sobre bytes ou sobre a palavra básica da maquina, enquanto operadores binários agem bit a bit.

Para efeito de exemplo considere seguintes variáveis: 

```golang
var v bin = 60 // 0011 1100
var b bin = 13 // 0000 1101
```

## AND

```golang
c := v & b
```

Equivalente a *0011-1100 & 0000-1101* e o resultado será *0000-1100*.

## OR

```golang
c = v | b
```
Equivalente a *0011-1100 | 0000-1101* e o resultado será *0011-1101*.

## NOT AND

```golang
c = v &^ b
```

Equivalente a *0011-1100 &^ 0000-1101* e o resultado será *0011-0000*. Uma curiosidade é que esse é o único operador que você realmente precisa, todos os outros operadores podem ser derivados de *NOT END*.

## LEFT SHIFT

```golang
c = v << 2
```

Equivalente a *0011-1100 << 0000-1101* e o resultado será *1111-0000*. Você também pode pensar no *LEFT SHIFT* como multiplicar por dois.

## RIGHT SHIFT

```golang
c = v >> 2
```

Equivalente a *0011-1100 >> 0000-1101* e o resultado será *0000-1111*. Você também pode pensar no *RIGHT SHIFT* como dividir por dois.

## XOR

```golang
c = v ^ b
```

Equivalente a *0011-1100 ^ 0000-1101* e o resultado será *0011-0001*. Eu sei que é tentador mas não use *XOR* para criptografia, pelo menos não sozinho.

