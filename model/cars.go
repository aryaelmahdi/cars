package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Car struct {
	Name         string `db:"name" json:"name"`
	Make         string `db:"make" json:"make"`
	Powerplant   string `db:"powerplant" json:"powerplant"`
	Aspiration   string `db:"aspiration" json:"aspiration"`
	Fuel         string `db:"fuel" json:"fuel"`
	Transmission string `db:"transmission" json:"transmission"`
	Drivetrain   string `db:"drivetrain" json:"drivetrain"`
	Type         string `db:"type" json:"type"`
	Image        []byte `db:"image"`
}

type CarModel struct {
	db *sqlx.DB
}

func (cm *CarModel) Init(db *sqlx.DB) {
	cm.db = db
}

func (cm *CarModel) InsertCar(newCar Car) *Car {
	query := "INSERT INTO cars (name, make, powerplant," +
		"aspiration, fuel, transmission, drivetrain, type, image)"
	if _, err := cm.db.Exec(query, &newCar.Name, &newCar.Make, &newCar.Powerplant, &newCar.Aspiration,
		&newCar.Fuel, &newCar.Transmission, &newCar.Transmission, &newCar.Drivetrain,
		&newCar.Type, &newCar.Image); err != nil {
		logrus.Error("Model : cannot insert car")
		return nil
	}
	return &newCar
}

func (cm *CarModel) GetAllCars() []Car {
	cars := []Car{}
	query := "SELECT * FROM cars"
	if err := cm.db.Select(&cars, query); err != nil {
		logrus.Error("Model : cannot get cars")
	}
	return cars
}

func (cm *CarModel) GetCarByName(name string) *Car {
	car := Car{}
	query := "SELECT * FROM cars where name = ?"
	if err := cm.db.Get(&car, query, &name); err != nil {
		logrus.Error("Model : cannot get car name")
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
