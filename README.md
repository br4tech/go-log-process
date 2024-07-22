# Go com Kafka: Adapter, Logs e Redpanda

Este projeto Go demonstra a integração com o Apache Kafka para publicação e consumo de mensagens, utilizando um tópico chamado "logs". 
Um adapter Kafka personalizado simplifica as operações, enquanto o Redpanda oferece uma interface visual para monitorar as mensagens. 
O Docker e o Docker Compose facilitam a configuração e execução do ambiente.

## Estrutura do Projeto

```bash
  |____README.md
  |____.gitignore
  |____Dockerfile
  |____cmd
  | |____main.go
  |____go.mod
  |____.vscode
  | |____launch.json
  |____.git
  |____go.sum
  |____docker-compose.yml
  |____internal
  | |____adapter
  | | |____kafka.go
  | |____port
  | | |____message.go
  | |____domain
  | | |____services
  | | | |____log.go
  | | |____entities
  | | | |____log.go
```

## Pré-requisitos

- Go (versão X.X ou superior)
- Docker
- Docker Compose

## Configuração

```bash
  git clone git@github.com:br4tech/go-log-process.git
```

### Inicie o Ambiente com Docker Compose:

```bash
  docker-compose up -d
```
