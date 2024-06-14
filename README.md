# Full Cycle - Go Export - Desafio Rate Limiter

O Rate Limiter e uma solução que visa limitar o número de requesições que podem ser realizadas por um determido Ip ou Token.

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

### 📋 Pré-requisitos

Usaremos o Docker para subir nossa imagem com o comando abaixo

```
docker-compose up --build
```

Para executar nossa aplicação devemos executar o comando abaixo, sendo o config.yaml o arquivo onde estará nossa configurações

```
go run cmd/main.go -config=config/config.yaml 
```

O serviço estara disponível o endereço abaixo

```
http://localhost:8080
```



