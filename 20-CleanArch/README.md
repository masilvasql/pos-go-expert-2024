### DESAFIO
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

### Como executar.

Executar o clone do projeto principal em sua máquina.

acessar a pasta 20-CleanArch e executar o comando:

```bash
docker-compose up -d 
```
________ 
### Para o GraphQL acessar a url: http://localhost:8080

Para testar a criação de uma order, no playground do GraphQL, executar a seguinte query:

```graphql
mutation createOrder{
  createOrder(input:{id:"dddsdasddd", Price:12.1, Tax:20.0}){
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```

Para listar as orders, no playground do GraphQL, executar a seguinte query:

```graphql
query listOrders{
  getOrders{
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```
_________
### Para o REST acessar a url: http://localhost:8000

Para testar a criação de uma order, acessar a pasta api onde terá o arquivo para listar as orders e criar uma nova order.

_______

### Para o GRPC acessar a url: http://localhost:50051

Para testar a criação e listagem, pode ser utilizado o evans como cliente GRPC.