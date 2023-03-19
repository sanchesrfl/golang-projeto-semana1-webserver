# Golang Webserver integrado ao MySQL

Repositório do mini-projeto da primeira semana de aula ministrada para a turma do DevInHouse no Lab365/Senai.

Objetivos:

- Apresentar a linguagem de maneira panorâmica orientado por projeto prático (webserver).

--- 


### Como executar o projeto.

#### Rodar MySQL no Docker localmente

requerimentos:

1. Docker (https://docs.docker.com/engine/install/)
2. MySQL CLI (https://dev.mysql.com/doc/mysql-shell/8.0/en/mysql-shell-install.html)

---
Com os requerimentos satisfeitos, execute os próximos passos:

- baixar imagem atual mysql

```bash
docker pull mysql:latest
```

- rodar mysql 

```bash
docker run --name mysql-db -e MYSQL_ROOT_PASSWORD=<password> -p 3306:3306 -d mysql:latest
```

- conecte no MySQL via CLI

```bash
mysql -u root -p<password> -h 127.0.0.1 -P 3306
```

- crie uma nova base de dados (schema)

```bash
CREATE DATABASE <nome>;
```

- selecione a database
USE <database>;

- para verificar se sua tabela foi criada posteriormente pelo webserver execute:

SHOW TABLES;

- para uma query simples validando se os dados foram registrados pelo webserver

SELECT * FROM <nome-da-table>

---

No arquivo /db/mysql.go insira os valores das variáveis na string de conexão

- nome da database (schema)
- password

---
Para executar o backend, abra seu terminal na pasta raiz do projeto e execute:

```bash
go run main.go
```

Para acessar o front-end e inserir dados na database
acesse em seu navegador:

localhost:8080