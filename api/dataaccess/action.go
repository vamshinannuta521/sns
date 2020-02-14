package dataaccess

import (
	"database/sql"
	"errors"

	"sns/models"
)

func (cl *PostgresClient) CreateAction(action *models.Action) error {
	query := ` INSERT INTO Action (event_name, action_type, action_spec, account_name) VALUES ($1, $2, $3, $4)`
	_, err := cl.DB.Exec(query, action.EventName, action.ActionType, action.ActionSpec, action.AccountName)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (cl *PostgresClient) GetAction(actionID string) (*models.Action, error) {

	query := ` SELECT id, event_name, action_type, action_spec, account_name FROM action where id = $1`
	rows, err := cl.DB.Query(query, actionID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	r, err := getModelFromDBEntitiesAction(rows)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("No row with given actionID")
	}
	return r[0], nil
}

func (cl *PostgresClient) GetAllActions() ([]*models.Action, error) {
	query := `SELECT id, event_name, action_type, action_spec, account_name FROM action`
	rows, err := cl.DB.Query(query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return getModelFromDBEntitiesAction(rows)
}

func (cl *PostgresClient) GetAllActionsWithEventFilter(eventName string) ([]*models.Action, error) {
	query := `SELECT id, event_name, action_type, action_spec, account_name FROM action where event_name = $1`
	rows, err := cl.DB.Query(query, eventName)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return getModelFromDBEntitiesAction(rows)
}

func getModelFromDBEntitiesAction(rows *sql.Rows) ([]*models.Action, error) {
	actions := make([]*models.Action, 0)

	defer rows.Close()
	for rows.Next() {
		var action models.Action
		err := rows.Scan(&action.ID, &action.EventName, &action.ActionType, &action.ActionSpec, &action.AccountName)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		actions = append(actions, &action)
	}

	return actions, nil
}
