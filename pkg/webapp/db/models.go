package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type TokenStore oauth2.Token

// User db model for users
type User struct {
	UserName string
	Email    string
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
