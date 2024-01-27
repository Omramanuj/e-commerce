package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    FullName  string
    Email     string `gorm:"unique"`
    UserType  string 
}

type Order struct {
    gorm.Model
    UserID     uint
    User       User
    OrderItems []OrderItem
    Status     string
    Total      float64
}

type OrderItem struct {
    gorm.Model
    OrderID   uint
    Order     Order
    ProductID uint
    Product   Product
    Quantity  int
    Price     float64
}

type Category struct {
    gorm.Model
    Name     string
    ParentID *uint
    Children []Category `gorm:"foreignKey:ParentID"`
    Products []Product  `gorm:"foreignKey:CategoryID"`
}

type Product struct {
    gorm.Model
    Name        string
    Description string
    Price       float64
    CategoryID  uint
    Category    Category
}
