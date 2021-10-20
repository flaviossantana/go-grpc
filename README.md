# GO gRPC

Olá! Nesse projeto vou aplicar um estudo de comunicação com gRPC, usando a linguagem GO Lang.



## gRCP
Documentação do gRPC fica em [grpc.io](https://grpc.io/docs/what-is-grpc/introduction/).

### A história por trás do gRPC
 
O gRPC foi inicialmente criado pelo Google, que usou uma única infraestrutura RPC de uso geral chamada Stubby para conectar o grande número de microsserviços em execução dentro e entre seus data centers por mais de uma década. 

### Visão geral
No gRPC, um aplicativo cliente pode chamar diretamente um método em um aplicativo de servidor em uma máquina diferente como se fosse um objeto local, facilitando a criação de aplicativos e serviços distribuídos. Como em muitos sistemas RPC, o gRPC é baseado na ideia de definir um serviço, especificando os métodos que podem ser chamados remotamente com seus parâmetros e tipos de retorno. No lado do servidor, o servidor implementa essa interface e executa um servidor gRPC para lidar com as chamadas do cliente. No lado do cliente, o cliente tem um stub que fornece os mesmos métodos que o servidor.

![alt text](https://grpc.io/img/landing-2.svg)


## Protocol Buffers
Documentação do Protocol Buffers
 fica em [developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers/docs/overview).

O gRPC usa Protocol buffers como meio de comunicação entre as requisições e respostas.


### O que são Protocol Buffers? 
Buffers de protocolo são o mecanismo extensível de linguagem neutra, plataforma-neutra e extensível do Google para serializar dados estruturados - pense em XML, mas menor, mais rápido e mais simples. Você define como deseja que seus dados sejam estruturados uma vez e, em seguida, pode usar o código-fonte gerado especial para escrever e ler facilmente seus dados estruturados de e para uma variedade de fluxos de dados e usando uma variedade de linguagens.

## Instalações


### Go lang
[Baixe e instale o GO Lang](https://golang.org/dl/).

Após instalação verifique a versão do GO:

```bash
~ % go version
go version go1.15 darwin/amd64
```

## Iniciando o Projeto

Criando o modulo do projeto:
```bash
% go mod init github.com/flaviossantana/go-grpc 
go: creating new go.mod: module github.com/flaviossantana/go-grpc

```
Instalando compilador protoc:
https://grpc.io/docs/protoc-installation/

MacOS, usando o Homebrew:
```bash 
$ brew install protobuf
$ protoc --version  # Ensure com piler version is 3+

% protoc --version
libprotoc 3.17.3
```

Para finalizar, temos que adicionar a pasta "/go/bin" no PATH para que tudo que seja instalado nesta pasta esteja disponível 
como comandos no terminal.
```bash
export PATH="$PATH:/Users/flaviosantana/go/bin"
```

Execute o comando abaixo para atualizar seu terminal:
```bash
% source ~/.zprofile
```

Compilando o meu arquivo .proto:
```bash
% protoc --proto_path=proto proto/*.proto --go_out=pb
```

Gerar meu stub de grpc:
```bash
% protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
 ```

Rodando o nosso servidor:
```bash 
 % go run cmd/server/server.go
```

Testando as chamadas gRPC via Evans:
https://github.com/ktr0731/evans#not-recommended-go-get
```bash 
 $ go get github.com/ktr0731/evans
 $ go install github.com/ktr0731/evans
 evans -r repl --host localhost --port 50051
 
   ______
  |  ____|
  | |__    __   __   __ _   _ __    ___
  |  __|   \ \ / /  / _. | | '_ \  / __|
  | |____   \ V /  | (_| | | | | | \__ \
  |______|   \_/    \__,_| |_| |_| |___/
 
  more expressive universal gRPC client
 
 pb.UserService@localhost:50051> service UserServicer 
 pb.UserService@localhost:50051> call AddUser
   id (TYPE_STRING) => 0
   name (TYPE_STRING) => Flavio Santana
   email (TYPE_STRING) => 9090falvio@gmai.com
   {
     "id": "2312",
     "email": "9090falvio@gmai.com",    
     "name": "Flavio Santana"
   }
 
 ```