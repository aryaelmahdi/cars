package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Fuel struct {
	Fuel string `json:"fuel" db:"fuel"`
}

type FuelModel struct {
	db *sqlx.DB
}

func (f *FuelModel) Init(db *sqlx.DB) {
	f.db = db
}

func (f *FuelModel) GetAllFuel() []Fuel {
	fuel := []Fuel{}
	query := "SELECT * FROM fuels"
	if err := f.db.Select(&fuel, query); err != nil {
		logrus.Error("Model : cannot get fuel")
	}
	return fuel
}

func (f *FuelModel) GetFuelByName(name string) *Fuel {
	fuel := Fuel{}
	query := "SELECT * from fuels where fuel = ?"
	if err := f.db.Get(&fuel, query, &name); err != nil {
		logrus.Error("Model : cannot get fuel name")
		return nil
	}
	return &fuel
}

func (f *FuelModel) DeleteFuel(name string) error {
	query := "DELETE FROM fuels where fuel = ?"
	if _, err := f.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete fuel")
		return err
	}
	return nil
}

func (f *FuelModel) InsertFuel(newFuel Fuel) *Fuel {
	query := "INSERT INTO fuels (fuel) VALUES (?)"
	_, err := f.db.Exec(query, newFuel.Fuel)
	if err != nil {
		logrus.Error("Model : cannot insert fuel")
		return nil
	}
	return &newFuel
}
