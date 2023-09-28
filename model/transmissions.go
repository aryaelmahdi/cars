package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Transmission struct {
	Type string `db:"type" json:"type"`
}

type TransmissionModel struct {
	db *sqlx.DB
}

func (tm *TransmissionModel) Init(db *sqlx.DB) {
	tm.db = db
}

func (tm *TransmissionModel) InsertTransmisson(newTrans Transmission) *Transmission {
	query := "INSERT INTO transmissions (type) VALUES (?)"
	if _, err := tm.db.Exec(query, &newTrans.Type); err != nil {
		logrus.Error("Model : cannot insert transmission")
		return nil
	}
	return &newTrans
}

func (tm *TransmissionModel) GetAllTransmissions() []Transmission {
	transmissions := []Transmission{}
	query := "SELECT * FROM transmissions"
	if err := tm.db.Select(&transmissions, query); err != nil {
		logrus.Error("Model : cannot get transmissions")
	}
	return transmissions
}

func (tm *TransmissionModel) GetTransmissionsByName(name string) *Transmission {
	transmission := Transmission{}
	query := "SELECT * FROM transmissions where name = ?"
	if err := tm.db.Get(&transmission, query, &name); err != nil {
		logrus.Error("Model : cannot get transmission name")
		return nil
	}
	return &transmission
}

func (tm *TransmissionModel) DeleteTransmission(name string) error {
	query := "DELETE from transmissions where name = ?"
	if _, err := tm.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete transmission")
		return err
	}
	return nil
}
