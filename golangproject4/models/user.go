package models

import "golangproject4/dao"

type User struct {
	Id       int
	Username string
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username)
}

func DeleteUser(id int) error {
	err := dao.Db.Where("id = ?", id).Delete(&User{}).Error
	return err
}

func GetUserList() ([]User, error) {
	var users []User
	err := dao.Db.Where("id < ?", 3).Find(&users).Error
	return users, err
}
