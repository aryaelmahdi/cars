package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Drivetrain struct {
	Drivetrain string `json:"drivetrain" db:"drivetrain"`
}

type DrivetrainModel struct {
	db *sqlx.DB
}

func (d *DrivetrainModel) Init(db *sqlx.DB) {
	d.db = db
}

func (d *DrivetrainModel) GetAllDrivetrain() []Drivetrain {
	drivetrains := []Drivetrain{}
	query := "SELECT * FROM drivetrains"
	if err := d.db.Select(&drivetrains, query); err != nil {
		logrus.Error("Model : cannot get drivetrains")
	}
	return drivetrains
}

func (d *DrivetrainModel) GetDrivetrainByName(name string) *Drivetrain {
	drivetrain := Drivetrain{}
	query := "SELECT * from drivetrains where drivetrain = ?"
	if err := d.db.Get(&drivetrain, query, &name); err != nil {
		logrus.Error("Model : cannot get drivetrain name")
		return nil
	}
	return &drivetrain
}

func (d *DrivetrainModel) DeletedriveTrain(name string) error {
	query := "DELETE FROM drivetrains where drivetrain = ?"
	if _, err := d.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete drivetrain")
		return err
	}
	return nil
}

func (d *DrivetrainModel) InsertDrivetrain(newDrive Drivetrain) *Drivetrain {
	query := "INSERT INTO drivetrains (drivetrain) VALUES (?)"
	_, err := d.db.Exec(query, newDrive.Drivetrain)
	if err != nil {
		logrus.Error("Model : cannot insert drivetrain")
		return nil
	}
	return &newDrive
}
