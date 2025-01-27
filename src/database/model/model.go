package model

import "gorm.io/gorm"

type User struct {
	Id        uint   `gorm:"primary_key; not null" json:"id"`
	Name      string `gorm:"not null; unique" json:"name"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash" gorm:"not null"`
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
	Value         int         `gorm:"check:Value <= 3" json:"value"`
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
