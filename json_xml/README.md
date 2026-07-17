# JSON e XML

JSON e XML continuam aparecendo no mesmo tipo de código, às vezes no mesmo serviço. API nova costuma falar JSON. Integração antiga, nota fiscal, SOAP perdido no canto e arquivo de configuração legado ainda chegam em XML.

O ponto deste exemplo é pequeno: uma struct `User`, tags nos campos e duas codificações saindo do mesmo valor. O campo `Email` usa `omitempty`; quando ele está vazio, some do JSON e do XML. Isso é útil, mas também perigoso: campo ausente e campo vazio viram coisas diferentes para quem recebe.

## Pré-requisitos

- Go 1.26 ou mais novo.
- Noção básica de struct e tags.

## Executar

```bash
go run .
```

Saída esperada:

```text
{
  "id": 7,
  "name": "Ada",
  "email": "ada@example.com"
}
<user>
  <id>7</id>
  <name>Ada</name>
  <email>ada@example.com</email>
</user>
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes travam dois comportamentos que dão bug fácil em integração: `DecodeJSON` preenche a struct como esperado e `omitempty` remove `email` quando ele está vazio.

## Pontos de atenção

- `encoding/json` ignora campos não exportados. Campo com letra minúscula não entra no JSON, mesmo com tag.
- `encoding/xml` precisa de um nome de elemento raiz; aqui ele vem de `XMLName xml.Name` com `xml:"user"`.
- `omitempty` não valida dado. Ele só decide se o campo entra na saída.

## O que fica fora

Este tópico não entra em schema, streaming com `Decoder`, namespaces XML nem validação de contrato. Isso merece exemplos separados; colocar tudo aqui só deixaria o primeiro contato pior.
