package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Repo ..
type User struct {
	gorm.Model
	Name       string    // 名字
	ApiKey     string    // api key
	LastTravel time.Time // 上次被遍历的时间
}

func (user *User) Create() error {
	return engine.Create(user).Error
}

func (user *User) UpdateByID() error {
	return engine.Model(new(User)).Where(user.ID).Updates(user).Error
}

func (user *User) FindByID() (u *User, err error) {
	u = new(User)
	err = engine.Where("id = ?", user.ID).Find(u).Error
	if err != nil {
		return
	}
	return
}

func (user *User) FindByName() (u *User, err error) {
	u = new(User)
	err = engine.Where("name = ?", user.Name).Find(u).Error
	if err != nil {
		return
	}
	return
}

func (user *User) FindAll() (users []*User, err error) {
	err = engine.Find(&users).Error
	if err != nil {
		return
	}
	return
}
