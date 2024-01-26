# RSV (Rows of String Values)

É uma alternativa binaria ao CSV. 

Um documento RSV representa um array de arrays de string que podem ser null.

```
array<array<string|null>>
```

O proposito é armazenar dados tabulares mas não esta limitado a isso, linhas podem comter qualquer quantidade de dados.

- Não contem tipos de dados, apenas strings e null, o proposito é tornar o formato simples e universal. A interpretação das strings é responsabilidade do programador.
- Como é um formato binario RSV não é feito para ser aberto/alterado por um editor de textos.
- Evita colisão de delimitadores usando UTF-8.
- RSV é codificado usando UTF-8 e os separadores são caracteres invalidos UTF-8.

EOV (End Of Value) caracter 255, 0xFF.
NULL Value caracter 254, 0xFE.
EOR (End Of Row) caracter 253, 0xFD.


-- TODO: Explicar caracter de terminação e caracter de separação (CSV usa caracteres de separação, RSV usa caracteres de terminação)

