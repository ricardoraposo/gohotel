package data

import (
	"fmt"

	"github.com/ricardoraposo/gohotel/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type UpdateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r UpdateUserRequest) ToBson() bson.M {
	m := bson.M{}
	if len(r.FirstName) > 0 {
        m["firstName"] = r.FirstName
	}
	if len(r.LastName) > 0 {
        m["lastName"] = r.LastName
	}
    return m
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r CreateUserRequest) Validate() map[string]string {
	errors := map[string]string{}
	if len(r.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("First name should have more than %d characters", minFirstNameLen)
	}
	if len(r.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("Last name should have more than %d characters", minLastNameLen)
	}
	if len(r.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("Password should have more than %d characters", minPasswordLen)
	}
	if !utils.IsValidEmail(r.Email) {
		errors["email"] = fmt.Sprintf("Email is not a valid one")
	}
	return nil
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassowrd string             `bson:"encryptedPassowrd" json:"-"`
}

func NewUserFromRequest(r CreateUserRequest) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	user := &User{
		FirstName:         r.FirstName,
		LastName:          r.LastName,
		Email:             r.Email,
		EncryptedPassowrd: string(encryptedPassword),
	}

	return user, nil
}
