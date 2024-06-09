### Go workspaces

 - Comando : go work init .\math\ .\sistema\
    go work ini: Cria um arquivo go.work no diretório atual
    .\math\ .\sistema\ : Cria um workspace com os diretórios math e sistema

- Rodando go mod tidy quando se tem workspaces ainda não publicados
    - go mod tidy: Atualiza o arquivo go.mod com as dependências do projeto
    - go mod tidy -e: Atualiza o arquivo go.mod com as dependências do projeto, mas não remove as que não estão sendo utilizadas
    