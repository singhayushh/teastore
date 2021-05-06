package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User schema
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id" form:"id"`
	Type      string    `gorm:"size:15;default:'Customer'" json:"type" form:"type"` // Either an Admin or Customer(by default)
	Name      string    `gorm:"size:255;not null;" json:"name" form:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" form:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password" form:"password"`
	Image     string    `gorm:"default:'https://raw.githubusercontent.com/Simulacra-Technologies/teastore/master/templates/profile.png'" json:"image" form:"image"`
	Address   string    `gorm:"default:'Blank'" json:"address" form:"address"`
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

// EncryptPassword uses the hash function to hash the user's password
func (user *User) EncryptPassword() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Validate function is a utility function that implements password hashing, updating the time entries etc and checks if entries are valid, based on actions being used for (login, update etc)
func (user *User) Validate(action string) error {

	user.Name = strings.TrimSpace(user.Name)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Image = html.EscapeString(strings.TrimSpace(user.Image))
	user.Address = strings.TrimSpace(user.Address)

	switch strings.ToLower(action) {
	case "login":
		if user.Email == "" {
			return errors.New("email is required")
		}
		if user.Password == "" {
			return errors.New("password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "update":
		if user.Name == "" {
			return errors.New("name is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		if user.Address == "" {
			return errors.New("address is required")
		}
		if user.Image == "" {
			user.Image = "https://raw.githubusercontent.com/Simulacra-Technologies/teastore/master/templates/profile.png"
		}
		return nil
	default:
		user.ID = 0 // gets auto set anyway
		if user.Name == "" {
			return errors.New("name is required")
		}
		if user.Address == "" {
			return errors.New("address is required")
		}
		if user.Email == "" {
			return errors.New("email is required")
		}
		if user.Password == "" {
			return errors.New("password is required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

// Save the user in the database
func (user *User) Save(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

// FetchAll returns an array of Users listing upto 100 users
func (user *User) FetchAll(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, err
}

// FetchByID needs a uint64 uid to search for the corresponding user
func (user *User) FetchByID(db *gorm.DB, uid uint64) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("User Not Found")
	}
	return user, err
}

// Update requires user auth. Allowed - name, email, image, address
func (user *User) Update(db *gorm.DB, uid uint64) (*User, error) {

	// To hash the password
	err := user.Validate("update")
	if err != nil {
		return nil, err
	}

	// Update the user
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":       user.Name,
			"email":      user.Email,
			"Image":      user.Image,
			"address":    user.Address,
			"updated_at": time.Now(),
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

// Delete is my favourite. Every user should use this.
func (user *User) Delete(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
