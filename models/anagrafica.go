package models

import (
	"gorm.io/gorm"
)

type Anagrafica struct {
	gorm.Model
	ID          int
	Name        string
	Sourname    string
	RagSoc      string
	CodFisc     string
	PIva        string
	IndirizzoID int
	Indirizzo   Indirizzo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Indirizzo struct {
	ID    int
	Via   string
	Citta string
	CAP   string
}

//create a Anagrafica
func CreateAnagrafica(db *gorm.DB, Anagrafica *Anagrafica) (err error) {
	err = db.Create(Anagrafica).Error
	if err != nil {
		return err
	}
	return nil
}

//get Anagraficas
func GetAnagrafiche(db *gorm.DB, Anagrafica *[]Anagrafica) (err error) {
	err = db.Order("id desc, name").Find(Anagrafica).Error
	if err != nil {
		return err
	}
	return nil
}

//get Anagrafica by id
func GetAnagrafica(db *gorm.DB, Anagrafica *Anagrafica, id string) (err error) {
	err = db.Where("id = ?", id).First(Anagrafica).Error
	if err != nil {
		return err
	}
	return nil
}

//update Anagrafica
func UpdateAnagrafica(db *gorm.DB, Anagrafica *Anagrafica) (err error) {
	db.Save(Anagrafica)
	return nil
}

//delete Anagrafica
func DeleteAnagrafica(db *gorm.DB, Anagrafica *Anagrafica, id string) (err error) {
	db.Where("id = ?", id).Delete(Anagrafica)
	return nil
}
