package dataaccess

import (
	"sns/models"
)

func CreateAccount(account *models.Account) error {
	pgClient := PostgresClient{}
	err := pgClient.OpenDb("abcd", "efgh")
	if err != nil {
		return err
	}
	defer pgClient.DB.close()
	sqlInsertEvent := ` INSERT INTO Account (id, name) VALUES ($1, $2)`
	_, err = pgClient.DB.Exec(sqlInsertEvent, account.ID, account.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetAccount(account_uuid string) (*models.Event, error) {
	pgClient := PostgresClient{}
	err := pgClient.OpenDb("abcd", "efgh")
	if err != nil {
		return err
	}
	defer pgClient.DB.close()
	sqlInsertEvent := ` SELECT id, name FROM account where id = $1`
	rows, err = pgClient.DB.Query(sqlInsertEvent, event_uuid)
	if err != nil {
		return nil, err
	}
	r, err := getModelFromDBEntities(rows)
	if err != nil{
		return nil,err
	}
	return r[0],nil
}

func GetAllAccounts() ([]*models.Account, error) {
	pgClient := PostgresClient{}
    err := pgClient.OpenDb("abcd", "efgh")
    if err != nil {
        return nil, err
    }
	defer pgClient.DB.close()
    sqlInsertEvent := `SELECT id, name FROM account`
    rows, err := pgClient.DB.Query(sqlInsertEvent)
    if err != nil {
        return nil, err
    }
    return getModelFromDBEntities(rows)
}


	
func getModelFromDBEntities(rows *sql.Rows) ([]*models.Account, error){
	accounts := make([]*models.Account, 0)

	defer rows.Close()
	for rows.Next() {
		account := models.Account{}
		err = rows.Scan(&account.Id, &account.Name, &account.CreatedBy)
		if err != nil { //dont return 1 err, return consolidated ones
			return nill, err
		}
		accounts.append(account)
	}

	return accounts, nil
}

