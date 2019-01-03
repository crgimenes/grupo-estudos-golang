# Exemplos com banco de dados

## Instalando sqlx

```console
go get github.com/jmoiron/sqlx
```

## SSL

Conectando PostgreSQL usando SSL

Criando o certificado de testes.

```console
openssl req -new -x509 -days 365 -nodes -text -out server.crt \
  -keyout server.key -subj "/CN=example.com"
```

Edite o arquivo postgresql.conf para ativar o uso da chave SSL.

```conf
ssl = on
ssl_cert_file = 'server.crt'
ssl_key_file = 'server.key'
```

Reinicie o serviço e teste a conexão com o seguinte comando

```console
psql "sslmode=require"
```

Na string de conexão use `slmode=verify-full` no lugar de `sslmode=disable`. Como seu certificado não é assinado por uma entidade certificadora você vai precisar marcar ele como confiavel no seu sistema operacional. Claro que em produção o certo é usar um certificado assinado.

Veja a [documentação do PostgreSQL](https://www.postgresql.org/docs/11/ssl-tcp.html) para ver mais exemplos de configuração.
