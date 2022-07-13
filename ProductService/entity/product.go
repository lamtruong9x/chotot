package entity

import "time"

type Product struct {
	Id          int       `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	ProductName string    `json:"product_name,omitempty" gorm:"type:varchar(255);not null;unique"`
	UserId      int       `json:"user_id,omitempty" gorm:"type:int"`
	CatId       string    `json:"cat_id,omitempty" gorm:"type:VARCHAR(10)"`
	TypeId      string    `json:"type_id,omitempty" gorm:"type:VARCHAR(10)"`
	Price       float64   `json:"price,omitempty" gorm:"type:float"`
	State       bool      `json:"state,omitempty" gorm:"type:boolean"`
	CreatedTime time.Time `json:"created_time,omitempty" gorm:"type:timestamp"`
	ExpiredTime time.Time `json:"expired_time,omitempty" gorm:"type:timestamp"`
	Address     string    `json:"address,omitempty" gorm:"type:varchar(255)"`
	Content     string    `json:"content,omitempty" gorm:"type:varchar(255)"`
}
