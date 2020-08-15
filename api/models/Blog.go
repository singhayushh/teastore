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

// Blog schema - CUD to admin only
type Blog struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id" form:"id"`
	Path      string    `gorm:"size:6;unique" json:"path" form:"path"`
	Title     string    `gorm:"size:255;not null;" json:"title" form:"title"`
	Cover     string    `gorm:"default:'https://covers.unsplash.com/photo-1523920290228-4f321a939b4c?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=756&q=80'" json:"cover" form:"cover"`
	Author    string    `gorm:"size:255;not null;" json:"author" form:"author"`
	Text      string    `gorm:"not null;" json:"text" form:"text"`
	Hits      uint64    `gorm:"default:0" json:"hits" form:"hits"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Validate is a utility function to format the blog according to schema
func (blog *Blog) Validate(action string) error {
	if action == "" {
		blog.ID = 0
		blog.Path = html.EscapeString(strings.TrimSpace(blog.Path))
		if blog.Path == "" {
			blog.Path = utils.GenerateTextHash(6)
		}
	}
	blog.Title = strings.TrimSpace(blog.Title)
	blog.Cover = strings.TrimSpace(blog.Cover)
	blog.Author = strings.TrimSpace(blog.Author)
	if blog.Title == "" {
		return errors.New("Title is required")
	}
	if blog.Cover == "" {
		blog.Cover = "https://raw.githubusercontent.com/Simulacra-Technologies/teastore/master/templates/Cover%20Not%20Available.png"
	}
	if blog.Author == "" {
		return errors.New("Author is required")
	}
	return nil
}

// Save the blog in the db
func (blog *Blog) Save(db *gorm.DB) (*Blog, error) {
	var err error
	err = db.Debug().Create(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// FetchAll returns an array of Blogs
func (blog *Blog) FetchAll(db *gorm.DB) (*[]Blog, error) {
	var err error
	blogs := []Blog{}
	err = db.Debug().Model(&Blog{}).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return &blogs, err
}

// FetchByID needs the path to search for the corresponding blog.
func (blog *Blog) FetchByID(db *gorm.DB, path string) (*Blog, error) {
	var err error
	err = db.Debug().Model(Blog{}).Where("path = ?", path).Take(&blog).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Blog Not Found")
	}
	return blog, err
}

// Update currently allows changing the cover, author, price, stock
func (blog *Blog) Update(db *gorm.DB, path string) (*Blog, error) {

	err := blog.Validate("update")
	if err != nil {
		log.Fatal(err)
	}

	// Update the blog
	db = db.Debug().Model(&Blog{}).Where("path = ?", path).Take(&Blog{}).UpdateColumns(
		map[string]interface{}{
			"title":      blog.Title,
			"cover":      blog.Cover,
			"author":     blog.Author,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	// Fetch the blog
	err = db.Debug().Model(&Blog{}).Where("path = ?", path).Take(&blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// Delete the blog from the database
func (blog *Blog) Delete(db *gorm.DB, path string) (int64, error) {

	db = db.Debug().Model(&Blog{}).Where("path = ?", path).Take(&Blog{}).Delete(&Blog{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
