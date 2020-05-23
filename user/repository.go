package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var connectionString = "root:rootpasswd@tcp(127.0.0.1:3306)/app?autoReconnect=true&useSSL=false"

func openConnection() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:rootpasswd@(localhost:3306)/app")
}

// Save - asdf
func Save(user User) (User, error) {
	db, err := openConnection()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	db.Create(&user)
	return user, nil
}

// FindAll - asdfa
func FindAll() ([]User, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	return users, nil
}

// FindByID - asdf
func FindByID(id string) (User, error) {
	db, err := openConnection()
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	var user User
	db.First(&user, id)
	return user, nil
}

// Delete - asdf
func Delete(id string) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Delete(User{}, "id = ?", id)
	return nil
}
