package dataaccess

import (
	"sns/models"
)

func CreateAccount(account *models.Account) error {
	pgClient := PostgresClient{}
	err := pgClient.OpenDb("abcd", "efgh")
	sqlInsertEvent := ` INSERT INTO Account (id, name) VALUES ($1, $2)`
	pgClient.DB.Exec(sqlInsertEvent, account.ID, account.Name)
	if err != nil {
		return err
	}
	return nil
}
