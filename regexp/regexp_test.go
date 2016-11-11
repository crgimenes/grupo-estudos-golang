package main

import "testing"

func TestPegaImagens(t *testing.T) {

	var conteudoHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Titulo</title>
</head>

<body>
<img src='imagem1.jpg'>
<img src="imagem2.jpg">
</body>

</html>`

	imagens := pegaImagens(conteudoHTML)

	if len(imagens) != 2 {
		t.Error("Erro imagens devia ter dois elementos retornou", len(imagens), imagens)
	}

	if imagens[0] != "imagem1.jpg" ||
		imagens[1] != "imagem2.jpg" {
		t.Error("Retorno inesperado em pegaImagens:",
			imagens[0],
			imagens[1],
			"era esperado imagem1.jpg imagem2.jpg")

	}

}
