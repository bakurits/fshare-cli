package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type TokenStore oauth2.Token

// User db model for users
type User struct {
	Email    string `gorm:"primaryKey"`
	Password string
	Token    TokenStore
}

// Scan scan value into TokenStore, implements sql.Scanner interface
func (t *TokenStore) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal TokenStore value:", value))
	}

	result := TokenStore{}
	err := json.Unmarshal(b, &result)
	*t = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (t TokenStore) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// PasswordRestoreRequest stores data about password restore requests
type PasswordRestoreRequest struct {
	Token       string
	Email       string
	RequestDate time.Time
}
