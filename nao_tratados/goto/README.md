# goto e labels

A instrução goto tem uma má fama que vem do tempo do BASIC quando era usada indiscriminadamente e acabava tornando o código impossível de ler. Em linguagens modernas entretanto é uma instrução perfeitamente válida e desde que usada com critério pode ajudar a tornar seu código mais limpo.

Alem de goto as instruções break e continue também aceitam labels, isso é muito util para quando por exemplo se quer sair de um *for* aninhado em outro *for* ou especificar para qual dos *fors* aninhados se quer fazer *continue*.

Go tem algumas regras para usar com goto, continue e break:

- Não se pode saltar para fora de um escopo de função
- Não se pode saltar sobre a declaração de uma variável
- Quando se usa *break* seguido de um label a instrução imediatamente após o label precisa ser um *for*, *switch*, ou *select*.
- Quando se usa *continue* seguido de um label a instrução imediatamente após o label precisa ser um *for*.
