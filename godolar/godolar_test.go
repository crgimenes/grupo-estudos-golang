package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var originalSite = `
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
        <script type="text/javascript" src="/ptax/dtagent_ICA_6000500061013.js" data-dtconfig="rid=RID_1298139846|rpid=817804523|domain=bcb.gov.br|lab=1|reportUrl=dynaTraceMonitor|agentUri=/ptax/dtagent_ICA_6000500061013.js"></script><link rel="stylesheet" type="text/css" href="/ptax_internet/ncss/style.css">
        <title></title>
    </head>
    <body>

        <div style="padding-left: 5%;">
            <img src="http://www4.bcb.gov.br/gifs/quadro-p.gif">&nbsp;Cotação de fechamento do dólar no dia 22/11/2017, Quarta-feira:
        </div>
        <br>

        <div style="padding-left: 10%;">
            <ul style="padding-left: 15px;">
                <li>Dólar-dos-EUA:</li>
            </ul>
            <table cellspacing="1" summary="Cotação de fechamento do Dólar americano">
                <tbody>
                    <tr class="fundoPadraoBEscuro3">
                        <th>Data</th>
                        <th>Taxa de Compra</th>
                        <th>Taxa de Venda</th>
                    </tr>
                    <tr class="fundoPadraoBClaro2">
                        <td align="CENTER">22/11/2017</td>
                        <td align="right">3,2555</td>
                        <td align="right">3,2561</td>
                    </tr>
                </tbody>
            </table>
            <br>
            <br>
        </div>

        <div align="justify" style="font-size: 0.8em">
            <img src="http://www.bcb.gov.br/img/BulletAzul2.gif">
            &nbsp;O Banco Central não assume qualquer responsabilidade pela não 
            simultaneidade ou falta das informações prestadas, assim como por 
            eventuais erros de paridades das moedas, ou qualquer outro, salvo 
            a paridade relativa ao dólar dos Estados Unidos da América em relação 
            ao Real. Igualmente, não se responsabiliza  pelos atrasos ou indisponibilidade 
            de serviços de telecomunicação, interrupção, falha ou pelas imprecisões no 
            fornecimento dos  serviços ou  informações. Não assume, também, responsabilidade 
            por qualquer perda ou dano oriundo de tais interrupções, atrasos, falhas ou 
            imperfeições, bem como pelo uso inadequado das  informações  contidas na transação.
        </div>

    </body>
</html>
`
var errData = ` <!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=ISO-8859-1">
        <script type="text/javascript" src="/ptax/dtagent_ICA_6000500061013.js" data-dtconfig="rid=RID_1298139846|rpid=817804523|domain=bcb.gov.br|lab=1|reportUrl=dynaTraceMonitor|agentUri=/ptax/dtagent_ICA_6000500061013.js"></script><link rel="stylesheet" type="text/css" href="/ptax_internet/ncss/style.css">
        <title></title>
    </head>
    <body>
        <div style="padding-left: 10%;">
            <ul style="padding-left: 15px;">
                <li>Dólar-dos-EUA:</li>
            </ul>
            <table cellspacing="1" summary="Cotação de fechamento do Dólar americano">
                <tbody>
                    <tr class="fundoPadraoBEscuro3">
                        <th>Data</th>
                        <th>Taxa de Compra</th>
                        <th>Taxa de Venda</th>
                    </tr>
                    <tr class="fundoPadraoBClaro2">
                        <td align="CENTER">22/11/2017</td>
                        <td align="right">3,2555</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </body>
</html>`
var errSite = `<html><head></head><body></body></html>`

func FakeResponse(site string) (contents []byte) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, site)
	}
	req := httptest.NewRequest("GET", "http://example.com/dolar", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	contents, _ = io.ReadAll(resp.Body)
	return
}

func TestResponseTratamento(t *testing.T) {

	for _, test := range []struct {
		// Struct que define os dados de entrada e saida necessarios para os testes
		site   string
		compra string
		venda  string
		err    error
	}{
		// Casos de teste para a função
		{originalSite, "3,2555", "3,2561", nil},
		{errData, "", "", errTratamento},
		{errSite, "", "", errTratamento},
	} {
		response := FakeResponse(test.site)
		compra, venda, err := TrataRequest(response)

		if compra != test.compra {
			t.Errorf("Esperava compra %v e obiteve %v\n", test.compra, compra)
		}
		if venda != test.venda {
			t.Errorf("Esperava venda %v e obiteve %v\n", test.venda, venda)
		}
		if err != test.err {
			t.Errorf("Esperava err %v e obiteve %v\n", test.err, err)
		}
	}

}

func TestBuscaRequest(t *testing.T) {
	_, err := BuscaRequest()
	if err != nil {
		t.Error(err)
	}

}
