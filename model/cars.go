package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Car struct {
	Name       string `db:"name" json:"name"`
	Make       string `db:"make" json:"make"`
	Powerplant string `db:"powerplant" json:"powerplant"`
	Drivetrain string `db:"drivetrain" json:"drivetrain"`
	Image      []byte `db:"image"`
}

type CompleteCar struct {
	Car
	Configuration string `db:"configuration" json:"configuration"`
	Displacement  string `db:"displacement" json:"displacement"`
	Crankshaft    string `db:"crankshaft" json:"crankshaft"`
	Horsepower    string `db:"horsepower" json:"horsepower"`
	Torque        string `db:"torque" json:"torque"`
}

type CarModel struct {
	db *sqlx.DB
}

func (cm *CarModel) Init(db *sqlx.DB) {
	cm.db = db
}

func (cm *CarModel) InsertCar(newCar Car) *Car {
	query := "INSERT INTO cars (name, make, powerplant,drivetrain, image) VALUES (?,?,?,?,?)"
	if _, err := cm.db.Exec(query, &newCar.Name, &newCar.Make, &newCar.Powerplant,
		&newCar.Drivetrain, &newCar.Image); err != nil {
		logrus.Error("Model : cannot insert car", err)
		return nil
	}
	return &newCar
}

func (cm *CarModel) GetAllCars() []CompleteCar {
	cars := []CompleteCar{}
	query := "SELECT cars.*, `engines`.configuration, `engines`.displacement," +
		"`engines`.crankshaft, `engines`.horsepower, `engines`.torque from cars " +
		"Cross join `engines` where engines.id = cars.powerplant"
	if err := cm.db.Select(&cars, query); err != nil {
		logrus.Error("Model : cannot get cars", err)
	}
	return cars
}

func (cm *CarModel) GetCarByName(name string) *CompleteCar {
	car := CompleteCar{}
	query := "SELECT cars.*, engines.configuration, engines.displacement," +
		"engines.crankshaft, engines.horsepower, engines.torque from cars " +
		"join `engines` on cars.powerplant = `engines`.id where cars.name = ?"
	if err := cm.db.Get(&car, query, &name); err != nil {
		logrus.Error("Model : cannot get car name", err)
		return nil
	}
	return &car
}

func (cm *CarModel) DeleteCar(name string) error {
	query := "DELETE from cars where name = ?"
	if _, err := cm.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete car")
		return err
	}
	return nil
}
