package user

import (
	"flag"
	"log"
	"os"

	// mysql dependency to be implicit
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var maxNbConcurrentGoroutines = flag.Int("maxNbConcurrentGoroutines", 20, "the number of goroutines that are allowed to run concurrently")
var concurrentGoroutines = make(chan struct{}, *maxNbConcurrentGoroutines)

func openConnection() (*gorm.DB, error) {
	concurrentGoroutines <- struct{}{}
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln(err)
	}
	return db, err
}

// Save - asdf
func Save(user User, responseChan chan func() (User, error)) {
	db, err := openConnection()
	if err != nil {
		responseChan <- (func() (User, error) { return User{}, err })
	}
	defer db.Close()
	db.Create(&user)
	<-concurrentGoroutines
	responseChan <- (func() (User, error) { return user, nil })
}

// FindAll - asdfa
func FindAll(responseChan chan func() ([]User, error)) {
	db, err := openConnection()
	if err != nil {
		responseChan <- (func() ([]User, error) { return nil, err })
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	<-concurrentGoroutines
	responseChan <- (func() ([]User, error) { return users, nil })
}

// FindByID - asdf
func FindByID(id string, responseChan chan func() (User, error)) {
	db, err := openConnection()
	if err != nil {
		responseChan <- (func() (User, error) { return User{}, err })
	}
	defer db.Close()
	var user User
	db.First(&user, id)
	<-concurrentGoroutines
	responseChan <- (func() (User, error) { return user, nil })
}

// Delete - asdf
func Delete(id string, responseChan chan error) {
	db, err := openConnection()
	if err != nil {
		responseChan <- err
	}
	defer db.Close()
	db.Delete(User{}, "id = ?", id)
	<-concurrentGoroutines
	responseChan <- nil
}
