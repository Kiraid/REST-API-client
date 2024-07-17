package models

import (
	"log"
	"time"
	"example.com/app/db"
	//"github.com/pelletier/go-toml/query"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserID      int64      `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)  //.prepare useful when inserting data
	if err != nil {
		log.Printf("Error preparing query: %v\n", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v\n", err)
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)  //.query useful when fetching data 
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventbyID(id int64) (*Event, error){
	query := "SELECT * FROM EVENTS WHERE id = ?"
	row := db.DB.QueryRow(query, id)  //gives us back one single row
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func(event Event) Update() error {
	query := `
	UPDATE events
	SET name=?, description=?, location=?, dateTime=?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err!= nil {
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
	
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

} 

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id= ? AND user_id= ?" 
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}