# Workspace

Chegamos a configurar o básico para trabalharmos com Go, mas caso queiramos entender melhor o nosso diretório onde trabalhamos,
o nosso **workspace** é preciso customizar/configurar mais algumas coisas.

O nosso workspace é nada mais nada menos que o diretório incluido na variável de ambiente GOPATH, em alguns casos o GOPATH ele já
vem direcionado por padrão para o mesmo diretório da GOROOT, no caso o nosso bin da linguagem.

Dado que nós temos nossa GOPATH apontada para algum diretório precisamos entender qual a estrutura de pastas 
em que o Go trabalha.

**bin** - Essa pasta possui os arquivos compilados

**pkg** - Essa pasta possui os arquivos compilados de acordo com o seu SO (Sistema Operacional)

**src** - Essa pasta possui todos seus códigos fontes de projetos Go ou demais linguagens ...
