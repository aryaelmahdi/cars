package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Manufacturers struct {
	Name      string `db:"name" json:"name"`
	Founder   string `db:"founder" json:"founder"`
	Country   string `db:"country" json:"country"`
	YearFound int    `db:"year_founded" json:"year_founded"`
}

type ManufacturersModel struct {
	db *sqlx.DB
}

func (m *ManufacturersModel) Init(db *sqlx.DB) {
	m.db = db
}

func (m *ManufacturersModel) GetAllManufacturers() []Manufacturers {
	manufacturers := []Manufacturers{}
	query := "SELECT * FROM manufacturers"
	if err := m.db.Select(&manufacturers, query); err != nil {
		logrus.Error("Model : cannot get manufacturers")
	}
	return manufacturers
}

func (m *ManufacturersModel) GetManufacturersByName(name string) *Manufacturers {
	manufacturer := Manufacturers{}
	query := "SELECT * FROM manufacturers where name = ?"
	if err := m.db.Get(&manufacturer, query, &name); err != nil {
		logrus.Error("Model : cannot get manufacturer id")
		return nil
	}
	return &manufacturer
}

func (m *ManufacturersModel) InsertManufacturer(newMan Manufacturers) *Manufacturers {
	query := "INSERT INTO manufacturers (name, founder, year_founded, country) VALUES (?,?,?,?)"
	if _, err := m.db.Exec(query, newMan.Name, newMan.Founder, newMan.YearFound, newMan.Country); err != nil {
		logrus.Error("Model : cannot insert manufacturer")
		return nil
	}
	return &newMan
}

func (m *ManufacturersModel) DeleteManufacturer(name string) error {
	query := "DELETE from manufacturers where name = ?"
	if _, err := m.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete manufacturer")
		return err
	}
	return nil
}
