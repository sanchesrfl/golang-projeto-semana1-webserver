package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertData(nome, cidade string) error {

	//insira o nome do schema e password na sring de conexão
	dbConnectionString := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/<schema>", "<password>")

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado ao Mysql com sucesso!")

	stmt, err := db.Prepare("INSERT INTO mytable (nome,cidade) VALUES (?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nome, cidade)
	if err != nil {
		return err
	}
	fmt.Println("Dados foram inseridos com sucesso!")

	defer stmt.Close()

	return nil

}
