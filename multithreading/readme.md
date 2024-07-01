### Rodar o go para identificar race condition
go run -race main.go


### rodar o apache bench

ab -n 100 -c 10 localhost:3003/

ab = apache bench
-n = numero de requests
-c = numero de concorrencia
