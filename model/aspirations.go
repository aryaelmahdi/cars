package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Aspirations struct {
	Id string `db:"id" json:"id"`
}

type AspirationsModel struct {
	db *sqlx.DB
}

func (a *AspirationsModel) Init(db *sqlx.DB) {
	a.db = db
}

func (a *AspirationsModel) InsertAspiration(newAsp Aspirations) (*Aspirations, error) {
	query := "INSERT INTO aspirations (id) VALUES (?)"
	_, err := a.db.Exec(query, &newAsp.Id)
	if err != nil {
		logrus.Error("Model : failed to insert aspirations")
		return nil, err
	}
	return &newAsp, nil
}

func (a *AspirationsModel) DeleteAspiration(id string) error {
	query := "DELETE from aspirations where id = ?"
	_, err := a.db.Exec(query, &id)
	if err != nil {
		logrus.Error("Model : failed to delete aspirations")
		return err
	}
	return nil
}

func (a *AspirationsModel) GetAllAspirations() ([]Aspirations, error) {
	aspirations := []Aspirations{}
	query := "SELECT * from aspirations"
	if err := a.db.Select(&aspirations, query); err != nil {
		logrus.Error("Model : failed to get aspirations")

	}
	return aspirations, nil
}

func (a *AspirationsModel) GetAspirationByID(id string) *Aspirations {
	aspiration := Aspirations{}
	query := "SELECT * from aspirations where id = ?"
	if err := a.db.Get(&aspiration, query, id); err != nil {
		logrus.Error("Model : failed to get aspiration")
		return nil
	}
	return &aspiration
}
