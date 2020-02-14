package dataaccess

import (
	"database/sql"
	"errors"

	"sns/models"
)

func (cl *PostgresClient) CreateAccount(account *models.Account) error {
	query := ` INSERT INTO Account (name) VALUES ($1)`
	_, err := cl.DB.Exec(query, account.Name)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (cl *PostgresClient) GetAccount(accountID string) (*models.Account, error) {

	query := ` SELECT id, name FROM account where id = $1`
	rows, err := cl.DB.Query(query, accountID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	r, err := getModelFromDBEntitiesAccount(rows)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("No row with given accountID")
	}
	return r[0], nil
}

func (cl *PostgresClient) GetAllAccounts() ([]*models.Account, error) {
	query := `SELECT id, name FROM account`
	rows, err := cl.DB.Query(query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return getModelFromDBEntitiesAccount(rows)
}

func getModelFromDBEntitiesAccount(rows *sql.Rows) ([]*models.Account, error) {
	accounts := make([]*models.Account, 0)

	defer rows.Close()
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Name)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}
