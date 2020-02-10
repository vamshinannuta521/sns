package dataaccess

import (
	"sns/models"
)

func CreateEvent(event *models.Event) error {
	pgClient := PostgresClient{}
	err := pgClient.OpenDb("abcd", "efgh")
	sqlInsertEvent := ` INSERT INTO event (id, name, createdBy) VALUES ($1, $2, $3)`
	pgClient.DB.Exec(sqlInsertEvent, event.ID, event.CreatedBy, event.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetEvent(event_uuid string) error {
	pgClient := PostgresClient{}
	err := pgClient.OpenDb("abcd", "efgh")
	sqlInsertEvent := ` SELECT id, name, createdBy FROM event where id = $1`
	pgClient.DB.Exec(sqlInsertEvent, event_uuid)
	if err != nil {
		return err
	}
	return nil
}
