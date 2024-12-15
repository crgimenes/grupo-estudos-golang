# Pasma no terminal.

Este exemplo mostra o efeito plasma usando códigos ANSI no terminal.

![plasma](plasma_no_terminal.gif)

A parte interessante é pré-processar as contas mais pesadas [`precalculateValues(r, c int)`](https://github.com/crgimenes/grupo-estudos-golang/blob/898b0dbac1d3709f5865062a27ba2a61b04a6e2e/plasma/main.go#L48) e [`precalculateColorStrings()`](https://github.com/crgimenes/grupo-estudos-golang/blob/898b0dbac1d3709f5865062a27ba2a61b04a6e2e/plasma/main.go#L78) concatenação de strings e cores, dessa forma o processamento do efeito fica consideravelmente mais rápido.

Outro ponto interessante é que o programa se ajusta ao tamanho do terminal, interceptando automaticamente o sinal `SIGWINCH`. A função [`updateTerminalSize()`](https://github.com/crgimenes/grupo-estudos-golang/blob/898b0dbac1d3709f5865062a27ba2a61b04a6e2e/plasma/main.go#L37) captura o novo tamanho do terminal e atualiza as variáveis.

O programa também interpreta `SIGINT`, assim, para terminar o programa basta pressionar control+c (^C).

Ainda existe muito espaço para otimização e melhorias de segurança de código, por exemplo, poderíamos parar de usar variáveis globais.
