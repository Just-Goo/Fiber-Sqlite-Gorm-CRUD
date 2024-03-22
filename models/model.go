package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
}

type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	ProductRefer uint    `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    uint    `json:"product_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	CreatedAt time.Time `json:"created_at"`
}
