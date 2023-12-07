package main

import (
	"biblioteca/pkg/config"
	"biblioteca/pkg/model"
	"biblioteca/pkg/operations"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatal("Não foi possível configurar o banco de dados:", err)
	}
	defer db.Close()
	clienteRepo := operations.NewClienteOperations(db)

	novoCliente := model.Cliente{
		Nome:          "João Silva",
		CPF:           "12345678910",
		DataCadastro:  "2023-01-01",
		UltimaReserva: "2023-01-10",
		Email:         sql.NullString{String: "joao@email.com", Valid: true},
	}
	err = clienteRepo.CriarCliente(novoCliente)
	if err != nil {
		log.Fatal("Erro ao criar cliente:", err)
	}

	// READ
	cliente, err := clienteRepo.ObterCliente(1)
	if err != nil {
		log.Fatal("Erro ao obter cliente:", err)
	}
	fmt.Printf("Cliente Obtido: %+v\n", cliente)

	// UPDATE
	cliente.Nome = "João Pereira"
	err = clienteRepo.AtualizarCliente(cliente)
	if err != nil {
		log.Fatal("Erro ao atualizar cliente:", err)
	}

	// DELETE
	err = clienteRepo.DeletarCliente(1)
	if err != nil {
		log.Fatal("Erro ao deletar cliente:", err)
	}
}
