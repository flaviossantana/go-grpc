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




