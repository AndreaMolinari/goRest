package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Name     string
	Email    string
	Password string
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Order("id desc, name").Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func Login(db *gorm.DB, User *User, loginRequest *User) (err error) {
	err = db.Where("email = ?", loginRequest.Email).Where("password = ?", loginRequest.Password).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
