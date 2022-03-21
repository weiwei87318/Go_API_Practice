package models

import (
	orm "api_practice/database"
)

type User struct {
	//gorm.Model
	ID  int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

var Users []User

//CRUD
//Create
func (user User) Insert() (id int, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//Read
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return 
}

//Update
func (user *User) Update(id int) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//Delete
func (user *User) Destroy(id int) (Result User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
