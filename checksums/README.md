# Checksums

Checksum serve para detectar corrupção acidental: arquivo truncado, byte trocado no download, cópia quebrada em disco. Não é assinatura, não é senha, não prova autoria. CRC é rápido e bom para erro bobo; para comparar conteúdo com menos chance de colisão, use um hash criptográfico como `SHA-256`.

O exemplo usa três funções da biblioteca padrão:

- `hash/crc32`, comum em formatos antigos, rede e validação simples;
- `hash/crc64`, útil quando você quer um CRC maior sem mudar de família;
- `crypto/sha256`, melhor quando o identificador do conteúdo precisa ser mais forte que um CRC.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Terminal no diretório `checksums`.

## Executar

```bash
go run .
```

Saída esperada:

```text
crc32  d57bb496
crc64  5ca2c9f6a56b90c2
sha256 5caa2109d58ce3fefe483e8b7176b80a00485ecbb3d92b18e48204a6ed4fe876
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes usam o payload fixo `invoice:42:paid`. Isso deixa o exemplo chato de propósito: se alguém trocar a tabela do CRC64, mudar o formato hexadecimal ou mexer no payload sem perceber, o teste quebra.

## Pontos de atenção

- CRC detecta erro acidental, mas não resiste a alteração maliciosa. Se a entrada vem de alguém que pode atacar o sistema, CRC sozinho é a ferramenta errada.
- `sha256.Sum256` retorna um array de 32 bytes. Para imprimir ou gravar em texto, formate em hexadecimal com `%x`.
- Não compare checksums de texto antes de decidir a codificação e as quebras de linha. `\n` e `\r\n` geram valores diferentes.

## Próximos passos

- Calcular o checksum lendo de um arquivo com `io.Copy` e `hash.Hash`.
- Comparar o custo entre CRC32, CRC64 e SHA-256 com benchmark.
- Usar `hmac` quando a verificação também precisa de uma chave compartilhada.
