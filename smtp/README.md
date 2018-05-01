# smtp

### Enviando Email!

Com grande elegância e poucas linhas, golang através pacote net/smtp fornece as funções necessárias para disparo de e-mails. 

No exemplo abaixo podemos ver o uso do método SendMail em ação:


```go

package main

import (
	"log"
	"net/smtp"
)

func main() {

	//Realiza o setup da autorização do servidor de SMTP. Não se esqueça de configuar seu Gmail SMTP server...
	//https://support.google.com/a/answer/176600?hl=en
	//https://support.google.com/accounts/answer/6010255?hl=en

	//Criamos um slice do tipo string do tamanho máximo de 1 para receber nosso e-mail destinatário.
	recipients := make([]string, 1)
	recipients[0] = "email@dodestinatario.com"

	//Veja mais parâmetros em: https://golang.org/pkg/net/smtp/#SendMail
	err := smtp.SendMail(
		/* endereço do servidor de SMTP */ "smtp.gmail.com:25",
		/* mecanismo de autenticação*/ smtp.PlainAuth("", "seuemail@gmail.com", "suasenhagmail", "smtp.gmail.com"),
		/* e-mail de origem */ "seuemail@gmail.com",
		/*Mensagem no RFC 822-style*/ recipients,
		/* Corpo da mensagem */ []byte("Subject:Olá!\n\n Olá Fulano. Tudo de bom com Go!"))
	if err != nil {
		log.Fatal(err)
	}

}

```

E é isso... simples assim!
