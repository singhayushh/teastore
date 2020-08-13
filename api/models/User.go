package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Type      string    `gorm:"size:15;default:'Customer'" json:"type"` // Either an Admin or Customer(by default)
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Address   string    `gorm:"default:'Blank'" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash - generates hash of the string passed as params and return the hashed password and err which is nil if no errors are thrown
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword - Verifies whether a password has been previously hashed and returns err, which is nil if its a success
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Validate function is a utility function that implements password hashing, updating the time entries etc and checks if entries are valid, based on actions being used for (login, update etc)
func (user *User) Validate(action string) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.ID = 0 // gets auto set anyway
	user.Type = strings.TrimSpace(user.Type)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Password = string(hashedPassword)
	user.Address = strings.TrimSpace(user.Address)

	switch strings.ToLower(action) {
	case "login":
		if user.Type == "" {
			return errors.New("API: Type is required")
		}
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if user.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "update":
		if user.Name == "" {
			return errors.New("Name is required")
		}
		if user.Address == "" {
			return errors.New("Address is required")
		}
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if user.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if user.Name == "" {
			return errors.New("Name is required")
		}
		if user.Address == "" {
			return errors.New("Address is required")
		}
		if user.Email == "" {
			return errors.New("Email is required")
		}
		if user.Password == "" {
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// Save the user in the database
func (user *User) Save(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindAllUsers returns an array of User type listing upto 100 users
func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, err
}

// FindUserByID needs a uint64 uid to search for the corresponding user
func (user *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("User Not Found")
	}
	return user, err
}

// UpdateAUser currently allows changing the name, email, password and address
func (user *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := user.Validate("update")
	if err != nil {
		log.Fatal(err)
	}

	// Update the user
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":      user.Name,
			"email":     user.Email,
			"password":  user.Password,
			"address":   user.Address,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	// Fetch the user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteAUser is my favourite. Every user should use this.
func (user *User) DeleteAUser(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
