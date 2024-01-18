package models

type Product struct {
	Product_id   int    `json:"product_id" gorm:"primary key;unique"`
	Product_name string `json:"product_name"`
	Body         string `json:"body"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique"  `
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null;unique"`
}
