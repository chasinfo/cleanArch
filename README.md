## Em um projeto pronto, adicionar a listagem de ordens
- listagem das ordens
- criar endpoint rest
- criar endpoint grpc
- criar uma query com GraphQL
 
#### Sobre o sistema 
- Go Versão: 1.21
- Docker 25
- docker-compose 1.29
- container docker utilizados:
  - Mysql 5.7
  - RabbitMq

#### Utilização 

baixar o projeto
```
git clone https://github.com/chasinfo/go-client-server.git
```

Iniciar o container: 
```
cd cleanArch
docker-compose up -d
cd cmd/ordersystem
go main.go wire_gen.go 
```

Fazer requisições utilizando algum client Rest, Ex. Postman.
```
Post:
Get:
```