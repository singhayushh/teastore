package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id" form:"id"`
	CartItems []CartItem `gorm:"foreignKey:CartId" json:"cart_items" form:"cart_items"`
	UserId    uint64     `gorm:"not null" json:"user_id" form:"user_id"`
}

type CartItem struct {
	CartId    uint64  `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint64  `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	UserId    uint64  `gorm:"default:null"`
}

func (cart *Cart) Save(db *gorm.DB) (*Cart, error) {
	err := db.Debug().Create(&cart).Error

	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (cart *Cart) FetchCart(db *gorm.DB, uid uint64) (*Cart, error) {
	err := db.Debug().Model(Cart{}).Preload("CartItems").Preload("CartItems.Product").First(&cart, "user_id = ?", uid).Error

	return cart, err
}

func (cart *Cart) AddtoCart(db *gorm.DB, uid uint64, pid uint64, qty int) error {
	product := Product{}
	fetchedProduct, err := product.FetchByID(db, pid)
	if err != nil {
		return db.Error
	}

	getCart, err := cart.FetchCart(db, uid)
	if err != nil {
		return db.Error
	}

	item := CartItem{
		ProductId: fetchedProduct.ID,
		Quantity:  qty,
		CartId:    getCart.ID,
		UserId:    uid,
	}

	getCart.CartItems = append(getCart.CartItems, item)

	fmt.Println(getCart.CartItems)

	db = db.Debug().Model(&Cart{}).Save(&getCart)

	if db.Error != nil {
		return db.Error
	}

	return nil
}
