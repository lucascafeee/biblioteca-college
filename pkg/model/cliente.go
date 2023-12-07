package model

import "database/sql"

type Cliente struct {
	ID            int
	Nome          string
	CPF           string
	DataCadastro  string
	UltimaReserva string
	Email         sql.NullString
}
