package repository

import (
	"database/sql"
	"fmt"
	"gin-gonic-leads-api/model"
)

type LeadRepository struct {
	connection *sql.DB
}

func NewLeadRepository(connection *sql.DB) LeadRepository {
	return LeadRepository{
		connection: connection,
	}
}

func (pr *LeadRepository) Create(lead model.Lead) (int, error) {

	var id int64
	// Prepare o comando INSERT
	query := "INSERT INTO leads (name, email) VALUES (?, ?)"
	result, err := pr.connection.Exec(query, lead.Name, lead.Email)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Obtenha o ID do registro inserido
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(id), nil
}
