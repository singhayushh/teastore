package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"teastore/api/utils"
	"time"

	"github.com/jinzhu/gorm"
)

// Product schema - CUD to admin only
type Product struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id" form:"id"`
	Path        string `gorm:"size:6;unique" json:"path" form:"path"`
	Name        string `gorm:"size:255;not null;" json:"name" form:"name"`
	Image       string `gorm:"default:'https://raw.githubusercontent.com/Simulacra-Technologies/teastore/master/templates/Image%20Not%20Available.png'" json:"image" form:"image"`
	Description string `gorm:"not null;" json:"description" form:"description"`
	Price       string `gorm:"not null;" json:"price" form:"price"`
	Stock       string `gorm:"default:'TRUE';" json:"stock" form:"stock"`
	Hits        uint64 `gorm:"default:0" json:"hits" form:"hits"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Validate is a utility function to format the product according to schema
func (product *Product) Validate(action string) error {
	var err error
	currentTime := time.Now()
	if action == "" {
		product.ID = 0
		product.Path = html.EscapeString(strings.TrimSpace(product.Path))
		if product.Path == "" {
			product.Path = utils.GenerateTextHash(6)
		}
	}
	product.Name = strings.TrimSpace(product.Name)
	product.Image = strings.TrimSpace(product.Image)
	product.Description = strings.TrimSpace(product.Description)
	product.Price = strings.TrimSpace(product.Price)
	product.CreatedAt = currentTime.Format("2006-01-02")
	product.UpdatedAt = currentTime.Format("2006-01-02")

	if err != nil {
		return err
	}

	if product.Name == "" {
		return errors.New("name is required")
	}
	if product.Image == "" {
		product.Image = "https://raw.githubusercontent.com/Simulacra-Technologies/teastore/master/templates/Image%20Not%20Available.png"
	}
	if product.Description == "" {
		return errors.New("description is required")
	}
	if product.Price == "" {
		return errors.New("price is required")
	}
	return nil
}

// Save the product in the db
func (product *Product) Save(db *gorm.DB) (*Product, error) {
	err := db.Debug().Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FetchAll returns an array of Products
func (product *Product) FetchAll(db *gorm.DB) (*[]Product, error) {
	products := []Product{}
	err := db.Debug().Model(&Product{}).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, err
}

// FetchByID needs the path to search for the corresponding product.
func (product *Product) FetchByID(db *gorm.DB, path string) (*Product, error) {
	err := db.Debug().Model(Product{}).Where("path = ?", path).Take(&product).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Product Not Found")
	}
	return product, err
}

// Update currently allows changing the image, description, price, stock
func (product *Product) Update(db *gorm.DB, id uint64) (*Product, error) {

	err := product.Validate("update")
	if err != nil {
		log.Fatal(err)
	}

	// Update the product
	db = db.Debug().Model(&Product{}).Where("id = ?", id).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"name":        product.Name,
			"path":        product.Path,
			"image":       product.Image,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	// Fetch the product
	err = db.Debug().Model(&Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Delete the product from the database
func (product *Product) Delete(db *gorm.DB, id uint64) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", id).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
