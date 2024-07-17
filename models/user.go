package models

import (
	"errors"
	"log"

	"example.com/app/db"
	"example.com/app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)  //.query useful when fetching data 
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


func (u *User) Save() error {
	query :=  "INSERT INTO users(email, password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	} 
	hashedpass ,err  := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hashedpass)
	if err != nil {
		return err }
	id, err := result.LastInsertId()
	u.ID = id
	return err	
}

func (u *User) ValidateCredential() error {
	 query := "SELECT id, password FROM users WHERE email = ?"
	 row := db.DB.QueryRow(query, u.Email)

	 var retrievedPassword string
	 err := row.Scan( &u.ID, &retrievedPassword)
	 if err != nil {
		return err
	 }
	 passwordisvalid := utils.CheckPassword(u.Password, retrievedPassword)

	 if !passwordisvalid {
		return errors.New("credential invalid")
	 }
	 return nil

}
