package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	Id            string `db:"id" json:"id"`
	Cylinder      string `db:"cylinder" json:"cylinder"`
	Configuration string `db:"configuration" json:"configuration"`
	Displacement  string `db:"displacement" json:"displacement"`
	Crankshaft    string `db:"crankshaft" json:"crankshaft"`
	Aspirations   string `db:"aspirations" json:"aspirations"`
	Horsepower    int    `db:"horsepower" json:"horsepower"`
	Torque        int    `db:"torque" json:"torque"`
}

type EngineModel struct {
	db *sqlx.DB
}

func (e *EngineModel) Init(db *sqlx.DB) {
	e.db = db
}

func (e *EngineModel) GetAllEngine() []Engine {
	engines := []Engine{}
	query := "SELECT * FROM engines"
	if err := e.db.Select(&engines, query); err != nil {
		logrus.Error("Model : cannot execute query")
	}
	return engines
}

func (e *EngineModel) GetEngineByID(id string) *Engine {
	engine := Engine{}
	query := "SELECT * FROM engines where id = ?"
	if err := e.db.Get(engine, query, &id); err != nil {
		logrus.Error("Model : cannot get engine data")
		return nil
	}
	return &engine
}

func (e *EngineModel) InsertEngine(newEngine Engine) *Engine {
	query := "INSERT INTO engines (id, cylinder,configuratin, capacity, crankshaft, aspirations," +
		"horsepower, torque) VALUES (?,?,?,?,?,?,?,?)"
	if _, err := e.db.Exec(query, newEngine.Id, newEngine.Displacement, newEngine.Configuration,
		newEngine.Crankshaft, newEngine.Aspirations, newEngine.Horsepower, newEngine.Torque); err != nil {
		logrus.Error("Model : cannot insert engine")
		return nil
	}
	return &newEngine
}

func (e *EngineModel) DeleteEngine(id string) error {
	query := "DELETE from engines where id = ?"
	if _, err := e.db.Exec(query, &id); err != nil {
		logrus.Error("Model : cannot delete engine")
		return err
	}
	return nil
}
