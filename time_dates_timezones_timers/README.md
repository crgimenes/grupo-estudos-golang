# Tempo, datas, timezone e timers

Tempo em Go parece simples até aparecer timezone. `time.Time` não guarda "horário de Brasília" como texto: ele guarda um instante e uma localização. Se você parseia uma data sem dizer a localização certa, o bug costuma aparecer só quando alguém compara horário local com UTC.

Este exemplo usa três peças pequenas da biblioteca padrão:

- `time.LoadLocation`, para carregar `America/Sao_Paulo`;
- `time.ParseInLocation`, para interpretar uma string sem offset usando aquela localização;
- `time.NewTimer`, para esperar um evento sem travar o programa com busy wait.

## Pré-requisitos

- Go instalado.
- Dados de timezone disponíveis no sistema. Em imagens Linux muito pequenas, como algumas bases Alpine antigas, `time.LoadLocation("America/Sao_Paulo")` pode falhar se o pacote `tzdata` não estiver instalado.

## Executar

```bash
go run .
```

Saída esperada nesta versão do exemplo:

```text
parsed: 2026-03-08T10:00:00-03:00
timer fired
```

O `-03:00` vem da localização carregada. A string de entrada era `2026-03-08 10:00`, sem offset explícito; quem deu o contexto foi `ParseInLocation`.

## Testar

```bash
go test ./...
```

Os testes verificam duas coisas: a data parseada fica associada a `America/Sao_Paulo`, e o timer dispara dentro de um limite curto. O limite do teste é maior que a duração do timer de propósito, porque scheduler não é cronômetro de laboratório.

## Pontos de atenção

- Prefira UTC para armazenar e comparar instantes. Converta para timezone local na borda: entrada, saída, relatório, tela.
- `time.Parse` usa UTC quando o layout não traz timezone. Para texto local, use `time.ParseInLocation`.
- Timer criado e abandonado deve ser parado com `Stop`. Neste exemplo o canal é sempre consumido, então não há cleanup interessante para mostrar.

## Próximos passos

- Trocar o exemplo para `time.Ticker` e parar o ticker com `Stop`.
- Comparar `time.Parse`, `time.ParseInLocation` e `time.Date` com a mesma data.
- Adicionar um caso com horário de verão em uma localização que ainda usa DST, para ver como a biblioteca resolve horários ambíguos.
