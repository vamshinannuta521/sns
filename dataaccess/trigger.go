package dataaccess

import (
	"database/sql"
	"errors"

	"sns/models"
)

func (cl *PostgresClient) CreateTrigger(trigger *models.Trigger) error {
	query := ` INSERT INTO Trigger (event_id, account_id, message) VALUES ($1, $2, $3)`
	_, err := cl.DB.Exec(query, trigger.EventID, trigger.AccountID, trigger.Message)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (cl *PostgresClient) GetTrigger(triggerID string) (*models.Trigger, error) {

	query := ` SELECT id, event_id,  account_id, message FROM Trigger where id = $1`
	rows, err := cl.DB.Query(query, triggerID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	r, err := getModelFromDBEntitiesTrigger(rows)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("No row with given triggerID")
	}
	return r[0], nil
}

func (cl *PostgresClient) GetAllTriggers() ([]*models.Trigger, error) {
	query := `SELECT id, event_id,  account_id, message FROM trigger`
	rows, err := cl.DB.Query(query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return getModelFromDBEntitiesTrigger(rows)
}

func getModelFromDBEntitiesTrigger(rows *sql.Rows) ([]*models.Trigger, error) {
	triggers := make([]*models.Trigger, 0)

	defer rows.Close()
	for rows.Next() {
		var trigger models.Trigger
		err := rows.Scan(&trigger.ID, &trigger.EventID, &trigger.AccountID, &trigger.Message)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		triggers = append(triggers, &trigger)
	}

	return triggers, nil
}
