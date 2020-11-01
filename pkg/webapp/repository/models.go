package repository

import "golang.org/x/oauth2"

// User db model for users
type User struct {
	UserName string
	Email    string
	Token    oauth2.Token
}
