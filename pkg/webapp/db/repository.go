package db

import (
	"gorm.io/driver/sqlite"
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository API for accessing database
type Repository interface {
	GetUser(string) (User, error)
	AddUser(User) error
	UpdateUser(User) error
	GetPasswordRestoreInfo(string) (PasswordRestoreRequest, error)
}

// NewRepository returns new repository object
func NewRepository(dialect, connectionString string) (Repository, error) {
	db := initGorm(dialect, connectionString)
	if db == nil {
		return nil, errors.New("error while initializing gorm object")
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) GetUser(email string) (User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return u, errors.Wrap(err, "error while getting token")
	}
	return u, nil
}

func (r *repository) UpdateUser(user User) error {
	if err := r.db.Updates(user).Error; err != nil {
		return errors.Wrap(err, "error while updating user")
	}
	return nil
}

func (r *repository) AddUser(user User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "error while adding new user")
	}
	return nil
}
func (r *repository) GetPasswordRestoreInfo(_ string) (PasswordRestoreRequest, error) {
	return PasswordRestoreRequest{}, nil
}

type repository struct {
	db *gorm.DB
}

func openDB(dialect string, connectionString string) (*gorm.DB, error) {
	switch dialect {
	case "sqlite3":
		return gorm.Open(sqlite.Open(connectionString), nil)
	case "mysql":
		return gorm.Open(mysql.Open(connectionString), nil)
	default:
		return nil, errors.New("unknown dialect")
	}
}

func initGorm(dialect string, connectionString string) *gorm.DB {
	db, err := openDB(dialect, connectionString)
	if err != nil {
		log.Println(err)
		return nil
	}

	db.Set("gorm:table_options", "charset=utf8")
	_ = db.AutoMigrate(User{})
	return db
}
