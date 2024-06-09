### Executando testes em go

- Comando : go test .
    - go test .: Roda todos os testes do projeto
    - go test -v .: Roda todos os testes do projeto e exibe o nome dos testes que estão sendo executados
    - go test -cover .: Roda todos os testes do projeto e exibe a cobertura de código
    - go test -coverprofile="coverage.out" .: Roda todos os testes do projeto e exibe a cobertura de código em um arquivo coverage.out
    - go tool cover -html="coverage.out": Abre o arquivo coverage.out em um navegador para visualizar a cobertura de código
    - go test -run TestNomeFuncao .: Roda apenas o teste com o nome TestNomeFuncao
    - go test -run TestNomeFuncao -v .: Roda apenas o teste com o nome TestNomeFuncao e exibe o nome dos testes que estão sendo executados
    - go test -run TestNomeFuncao -v -cover .: Roda apenas o teste com o nome TestNomeFuncao e exibe a cobertura de código
    - go test -run TestNomeFuncao -v -coverprofile=coverage.out .: Roda apenas o teste com o nome TestNomeFuncao e exibe a cobertura de código em um arquivo coverage.out
    - go tool cover -html=coverage.out: Abre o arquivo coverage.out em um navegador para visualizar a cobertura de código

### Executando banchmark em go
    - Comando : go test -bench .
        - go test -bench .: Roda todos os benchmarks do projeto
        - go test -bench . -benchmem .: Roda todos os benchmarks do projeto e exibe a alocação de memória
        - go test -bench . -benchmem -run=^$: Roda todos os benchmarks do projeto e exibe a alocação de memória
        - go test -bench . -benchmem -run=^$ -v .: Roda todos os benchmarks do projeto e exibe a alocação de memória
        - go test -bench=NomeFuncao .: Roda apenas o benchmark com o nome NomeFuncao
        - go test -bench=NomeFuncao -benchmem .: Roda apenas o benchmark com o nome NomeFuncao e exibe a alocação de memória
        - go test -bench=NomeFuncao -benchmem -run=^$ .: Roda apenas o benchmark com o nome NomeFuncao e exibe a alocação de memória
        - go test -bench=NomeFuncao -benchmem -run=^$ -v .: Roda apenas o benchmark com o nome NomeFuncao e exibe a alocação de memória
        - go test -bench . -benchmem -count=10: Roda todos os benchmarks do projeto 10 vezes e exibe a alocação de memória
        -go test -bench . -benchmem -count=10 benchtime=10s: Roda todos os benchmarks do projeto 10 vezes e exibe a alocação de memória com um tempo de execução de 10s

### Executando fuzzing em go
    - Comando : 
        - go test -fuzz=.
        - go test -fuzz=FuzzNomeFuncao .
        - go test -fuzz=FuzzNomeFuncao .: Roda o fuzzing na função com o nome FuzzNomeFuncao
        - go test -fuzz=FuzzNomeFuncao -v .: Roda o fuzzing na função com o nome FuzzNomeFuncao e exibe o nome dos testes que estão sendo executados
        - go test -fuzz=FuzzNomeFuncao -count=10 .: Roda o fuzzing na função com o nome FuzzNomeFuncao 10 vezes
        - go test -fuzz=FuzzNomeFuncao -count=10 -v .: Roda o fuzzing na função com o nome FuzzNomeFuncao 10 vezes e exibe o nome dos testes que estão sendo 
        executados
        - go test -fuzz=. -run=^$: Roda o fuzzing em todas as funções do projeto
<!-- gerar arquivo  de resultado do erro -->
    - go test -fuzz=FuzzNomeFuncao -fuzzcount=1000000 -fuzztime=10s -fuzzout=crashers.txt .: Roda o fuzzing na função com o nome FuzzNomeFuncao 1000000 vezes com um tempo de execução de 10s e salva os resultados em um arquivo crashers.txt
    - go test -fuzz=FuzzNomeFuncao -fuzzcount=1000000 -fuzztime=10s -fuzzout=crashers.txt -v .: Roda o fuzzing na função com o nome FuzzNomeFuncao 1000000 vezes com um tempo de execução de 10s e salva os resultados em um arquivo crashers.txt e exibe o nome dos testes que estão sendo executados
### Ajuda
    - go help test