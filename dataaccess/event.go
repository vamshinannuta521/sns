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

func GetAllEvents() ([]*models.Event, error) {
	pgClient := PostgresClient{}
    err := pgClient.OpenDb("abcd", "efgh")
    if err != nil {
        return nil, err
    }
    sqlInsertEvent := `SELECT id, name, createdBy FROM event`
    rows, err := pgClient.DB.Query(sqlInsertEvent)
    if err != nil {
        return nil, err
    }
    return getModelFromDBEntities(rows)
}


	
func getModelFromDBEntities(rows *sql.Rows) ([]*models.Event, error){
	events := make([]*models.Event, 0)

	defer rows.Close()
	for rows.Next() {
		event := models.Event{}
		err = rows.Scan(&event.Id, &event.Name, &event.CreatedBy)
		if err != nil {
			return nill, err
		}
		events.append(event)
	}

	return events, nil
}
