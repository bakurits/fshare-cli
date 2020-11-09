package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

// Repository API for accessing database
type Repository interface {
	GetToken(string) (User, error)
}

// NewRepository returns new repository object
func NewRepository(connectionString string) (Repository, error) {
	db := initGorm("sqlite3", connectionString)
	if db == nil {
		return nil, errors.New("error while initializing gorm object")
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) GetToken(userName string) (User, error) {
	var u User
	if err := r.db.Where("user_name = ?", userName).First(&u).Error; err != nil {
		return u, errors.Wrap(err, "error while getting token")
	}
	return u, nil
}

func (r *repository) AddUser(user User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "error while adding new user")
	}
	return nil
}

type repository struct {
	db *gorm.DB
}

func initGorm(dialect string, connectionString string) *gorm.DB {
	var db, err = gorm.Open(dialect, connectionString)
	if err != nil {
		log.Println(err)
		return nil
	}

	db.Set("gorm:table_options", "charset=utf8")
	db.AutoMigrate(User{})
	return db
}
