package model

import "gorm.io/gorm"

type User struct {
	Id       uint   `gorm:"primary_key; not null" json:"id"`
	Name     string `gorm:"not null; unique" json:"name"`
	Email    string `gorm:"unique ;not null" json:"email"`
	Password string `json:"password"`
}

type ProgressClicker struct {
	Id       uint   `gorm:"primary_key; not null" json:"id"`
	UserId   uint   `gorm:"not null; index" json:"user_id"`
	User     User   `gorm:"foreignkey:UserId; references:Id" json:"-"`
	Clicks   int    `json:"clicks"`
	UserName string `json:"user_name"`
}

type Improvement struct {
	Id          uint   `gorm:"primary_key; not null" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserImprovement struct {
	Id            uint        `gorm:"primary_key; not null" json:"id"`
	UserId        uint        `gorm:"not null; index" json:"user_id"`
	UserName      string      `json:"user_name"`
	ImprovementId uint        `gorm:"not null; index" json:"improvement_id"`
	Improvement   Improvement `gorm:"foreignkey:ImprovementId; references:Id" json:"-"`
	Value         int         `json:"value"`
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	pc := &ProgressClicker{
		UserId:   u.Id,
		Clicks:   0,
		UserName: u.Name,
	}
	err := tx.Create(pc).Error
	if err != nil {
		return err
	}
	return nil
}
