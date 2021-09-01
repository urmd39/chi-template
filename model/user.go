package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Role string `json:"role" bson:"role"` // admin, driver, customer, superadmin
	jwt.StandardClaims
}
