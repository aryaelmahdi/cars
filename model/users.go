package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Users struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string
	Password string
}

type UserModel struct {
	db *sqlx.DB
}

func (u *UserModel) Init(db *sqlx.DB) {
	u.db = db
}

func (u *UserModel) Login(email string, password string) *Users {
	var userData = Users{}
	query := "Select * from users where email = ? and password = ?"
	if err := u.db.Get(&userData, query, email, password); err != nil {
		if err == sql.ErrNoRows {
			logrus.Error("Model : User not found", err)
		}
		logrus.Error("Model : Login failed,", err.Error())
		return nil
	}

	return &userData
}

func (u *UserModel) Register(newUser Users) (*Users, error) {
	var query = "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := u.db.Exec(query, &newUser.Name, &newUser.Email, &newUser.Password)
	if err != nil {
		logrus.Error("Failed to insert data:", err.Error())
		return nil, err
	}
	return &newUser, nil
}

func (u *UserModel) GetUserByID(id int) (*Users, error) {
	user := Users{}
	query := "SELECT name, email from users where id = ?"
	if err := u.db.Get(&user, query, &id); err != nil {
		logrus.Error("Model : failed to get data")
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) GetAllUsers() ([]Users, error) {
	var users []Users
	query := "SELECT name, email from users"
	if err := u.db.Select(&users, query); err != nil {
		logrus.Error("Model : failed to get data")
		return nil, err
	}
	return users, nil
}

func (u *UserModel) DeleteUser(id int) {
	var query = "DELETE from users where id = ?"
	_, err := u.db.Exec(query, id)
	if err != nil {
		logrus.Error("Model : failed to delete data")
	}
}
