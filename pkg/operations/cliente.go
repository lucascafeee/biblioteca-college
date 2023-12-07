package operations

import (
	"biblioteca/pkg/model"
	"database/sql"
)

type ClienteOperations struct {
	db *sql.DB
}

func NewClienteOperations(db *sql.DB) *ClienteOperations {
	return &ClienteOperations{db: db}
}

func (o *ClienteOperations) CriarCliente(c model.Cliente) error {
	query := `INSERT INTO biblioteca.clientes(nome, cpf, data_cadastro, ultima_reserva, email) VALUES(?, ?, ?, ?, ?)`
	_, err := o.db.Exec(query, c.Nome, c.CPF, c.DataCadastro, c.UltimaReserva, c.Email)
	return err
}

func (o *ClienteOperations) ObterCliente(id int) (model.Cliente, error) {
	var c model.Cliente
	query := `SELECT idclientes, nome, cpf, data_cadastro, ultima_reserva, email FROM biblioteca.clientes WHERE idclientes = ?`
	row := o.db.QueryRow(query, id)
	err := row.Scan(&c.ID, &c.Nome, &c.CPF, &c.DataCadastro, &c.UltimaReserva, &c.Email)
	return c, err
}

func (o *ClienteOperations) AtualizarCliente(c model.Cliente) error {
	query := `UPDATE biblioteca.clientes SET nome = ?, cpf = ?, data_cadastro = ?, ultima_reserva = ?, email = ? WHERE idclientes = ?`
	_, err := o.db.Exec(query, c.Nome, c.CPF, c.DataCadastro, c.UltimaReserva, c.Email, c.ID)
	return err
}

func (o *ClienteOperations) DeletarCliente(id int) error {
	query := `DELETE FROM biblioteca.clientes WHERE idclientes = ?`
	_, err := o.db.Exec(query, id)
	return err
}
