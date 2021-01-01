package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"user-auth/hash"
)

type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}

const userPwPepper = "secret-random-string"
const secretKey = "secret-random-string"

var RecordNotFound = errors.New("please create the account first")
var ErrInvalidPassword = errors.New("enter a valid password")

func NewUserService(connectionInfo string) (*UserService, error) {

	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(secretKey)
	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil

}

func (us *UserService) Create(user *User) error {

	pwBytes := []byte(user.Password + userPwPepper)

	HashedPass, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = ""
	user.PasswordHash = string(HashedPass)

	return us.db.Create(user).Error

}

// This is used to login a user
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}

	return foundUser, nil
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)

	if err != nil {
		return nil, err
	}
	return &user, err

}
func first(db *gorm.DB, dst interface{}) error {

	err := db.First(dst).Error

	if err == gorm.ErrRecordNotFound {
		return RecordNotFound
	}
	return err

}

func (us *UserService) DestructiveReset() error {
	err := us.db.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}

	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error {
	err := us.db.Migrator().AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)
	}
	return us.db.Model(&User{}).Where("email = ?", user.Email).Update("remember_hash", user.RememberHash).Error
	//return us.db.Save(user).Error
}
func (us *UserService) ByRememberToken(token string) (*User, error) {

	var user User

	tokenHash := us.hmac.Hash(token)
	db := us.db.Where("remember_hash =? ", tokenHash)
	err := first(db, &user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
