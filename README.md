# Golang Webserver integrado ao MySQL

Repositório do mini-projeto da primeira semana de aula ministrada para a turma do DevInHouse no Lab365/Senai.

Objetivos:

- Apresentar a linguagem de maneira panorâmica orientado por projeto prático.

--- 


### Como executar o projeto.

#### Rodar MySQL no Docker localmente.

Requerimentos:

1. Docker (https://docs.docker.com/engine/install/)
2. MySQL CLI (https://dev.mysql.com/doc/mysql-shell/8.0/en/mysql-shell-install.html)

---
Com os requerimentos satisfeitos, execute os próximos passos:

- Baixar imagem atual mysql

```bash
docker pull mysql:latest
```

- Rodar mysql 

```bash
docker run --name mysql-db -e MYSQL_ROOT_PASSWORD=<password> -p 3306:3306 -d mysql:latest
```

- Conecte no MySQL via CLI

```bash
mysql -u root -p<password> -h 127.0.0.1 -P 3306
```

- Crie uma nova base de dados (schema)

```bash
CREATE DATABASE <nome>;
```

- Selecione a database

```bash
USE <database>;
```

- Crie a tabela de users;

```bash
CREATE TABLE IF NOT EXISTS mytable (nome VARCHAR(255), cidade VARCHAR(255));
```

- Para verificar se sua tabela foi criada execute:

```bash
SHOW TABLES;
```


- Para uma query simples validando se os dados foram registrados pelo webserver

```bash
SELECT * FROM <nome-da-table>
```


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
