package dataaccess

import (
	"database/sql"
	"errors"
	"sns/models"
)

func (cl *PostgresClient) CreateEvent(event *models.Event) error {
	query := ` INSERT INTO event (name, account_name) VALUES ($1, $2)`
	_, err := cl.DB.Exec(query, event.Name, event.AccountName)
	if err != nil {
		logger.Errorf("Error in create event", err.Error())
		return err
	}
	logger.Infof("Successfully created event")
	return nil

}

func (cl *PostgresClient) GetEvent(eventID string) (*models.Event, error) {

	query := ` SELECT id, name, account_name FROM event where id = $1`
	rows, err := cl.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	r, err := getModelFromDBEntitiesEvent(rows)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("No row with given eventID")
	}
	return r[0], nil
}

func (cl *PostgresClient) GetAllEvents() ([]*models.Event, error) {
	query := `SELECT id, name, account_name FROM event`
	rows, err := cl.DB.Query(query)
	if err != nil {
		return nil, err
	}
	return getModelFromDBEntitiesEvent(rows)

}

func getModelFromDBEntitiesEvent(rows *sql.Rows) ([]*models.Event, error) {
	events := make([]*models.Event, 0)

	defer rows.Close()
	for rows.Next() {
		event := models.Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.AccountName)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
