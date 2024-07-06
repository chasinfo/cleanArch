## Em um projeto pronto, adicionar a listagem de ordens
- listagem das ordens
  - criar endpoint rest
  - criar endpoint grpc
  - criar endpoint com GraphQL
 
### Sobre o sistema 
- Go Versão: 1.21
- Docker 25
- docker-compose 1.29
- container docker utilizados:
  - Mysql 5.7
  - RabbitMq

### Utilização 

#### baixar o projeto
```bash
$ git clone https://github.com/chasinfo/cleanArch.git
```

#### subir os containers: 
```bash
$ cd cleanArch
$ docker-compose up -d
```

#### iniciar a aplicação
```bash
$ cd cmd/ordersystem
$ go main.go wire_gen.go 
```

####  Fazer requisições utilizando um client Rest, Ex. Postman.

#### Adicionar uma order
```
- Post: localhost:8000/order
  - body json:
```
```yml
{
    "id":"ssaaa",
    "price": 100.5,
    "tax": 0.5
}
```
#### Listar uma order
```bash
- Get: localhost:8000/order/list
```

#### Fazer requisições utilizando um client GraphQL

http://localhost:8080

```yml
mutation createOrder {
  createOrder(input: {id: "jujuba", Price: 13.5, Tax:3.0}) {
  	id
  	Price
  	Tax
  	FinalPrice
  }
}

query queryOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Fazer requisições utilizando GRPC

```bash
$ evans -r repl
```

#### listar ordens
```evans
127.0.0.1:50051> package pb

pb@127.0.0.1:50051> service OrderService

pb.OrderService@127.0.0.1:50051> call ListOrders
```