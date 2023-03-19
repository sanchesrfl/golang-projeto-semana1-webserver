package main

import (
	"fmt"
	"log"
	"net/http"
	"nomealeatorio/db"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {

		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request foi bem sucedido!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	//insere dados de input do usuário na database MySQL
	db.InsertData(name, address)

	fmt.Fprintf(w, "Nome = %s\n", name)
	fmt.Fprintf(w, "Cidade = %s\n", address)

}

//funcao principal do script

func main() {

	//servidor de arquivos estáticos do webserver
	fileServer := http.FileServer(http.Dir("./static"))

	//root handle
	http.Handle("/", fileServer)

	//form handle
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Iniciando web-servidor, na porta 8080.\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {

		log.Fatal(err)

	}

}
