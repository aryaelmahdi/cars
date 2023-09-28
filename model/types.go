package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Types struct {
	Name string `db:"name" json:"name"`
}

type TypesModel struct {
	db *sqlx.DB
}

func (t *TypesModel) Init(db *sqlx.DB) {
	t.db = db
}

func (t *TypesModel) InsertTypes(newTypes Types) *Types {
	query := "INSERT INTO types (name) VALUES (?)"
	if _, err := t.db.Exec(query, &newTypes.Name); err != nil {
		logrus.Error("Model : cannot insert types")
		return nil
	}
	return &newTypes
}

func (t *TypesModel) GetAllTypes() []Types {
	types := []Types{}
	query := "SELECT * FROM types"
	if err := t.db.Select(&types, query); err != nil {
		logrus.Error("Model : cannot get Types")
	}
	return types
}

func (t *TypesModel) GetTypeByName(name string) *Types {
	types := Types{}
	query := "SELECT * FROM types where name = ?"
	if err := t.db.Get(&types, query, &name); err != nil {
		logrus.Error("Model : cannot get type name")
		return nil
	}
	return &types
}

func (t *TypesModel) DeleteType(name string) error {
	query := "DELETE from types where name = ?"
	if _, err := t.db.Exec(query, &name); err != nil {
		logrus.Error("Model : cannot delete types")
		return err
	}
	return nil
}
