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

	var id int
	query, err := pr.connection.Prepare("INSERT INTO leads (name, email) VALUES ($1,$2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(lead.Name, lead.Email).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
