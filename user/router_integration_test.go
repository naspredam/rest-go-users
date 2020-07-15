//+build integration_tests

package user

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestGetAllUsersIntegration(t *testing.T) {
	db, err := stablishConnection()
	tx := db.Begin()
	defer db.Close()
	user1 := User{ID: 1, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	db.Create(&user1)
	user2 := User{ID: 2, FirstName: "xyz", LastName: "pqr", Phone: "1234567890"}
	db.Create(&user2)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Router()

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","phone":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","phone":"1234567890"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	tx.Rollback()
}

func TestCreateNewUserIntegration(t *testing.T) {
	db, err := stablishConnection()
	tx := db.Begin()
	defer db.Close()

	jsonStr := []byte(`{"first_name":"Krish","last_name":"Bhanushali","phone":"0987654321"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Router()

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		return
	}

	// Check the response body is what we expect.
	expected := `{"id":[0-9]+,"first_name":"Krish","last_name":"Bhanushali","phone":"0987654321"}`
	r, _ := regexp.Compile(expected)
	if !r.MatchString(rr.Body.String()) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	tx.Rollback()
}

func TestFetchUserByIdIntegration(t *testing.T) {
	db, err := stablishConnection()
	tx := db.Begin()
	defer db.Close()
	userFromRepository := User{ID: 12, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	db.Create(&userFromRepository)

	req, err := http.NewRequest("GET", "/users/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Router()

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}

	// Check the response body is what we expect.
	expected := `{"id":12,"first_name":"Krish","last_name":"Bhanushali","phone":"0987654321"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	tx.Rollback()
}

func TestDeleteUserByIdIntegration(t *testing.T) {
	id := 22
	db, err := stablishConnection()
	tx := db.Begin()
	defer db.Close()
	userFromRepository := User{ID: id, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	db.Create(&userFromRepository)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := Router()

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		return
	}

	if rr.Body.String() != "" {
		t.Errorf("body must be nil")
	}

	var user *User
	db.First(user, id)
	if user != nil {
		t.Errorf("the user %v should have been deleted from the database", id)
		return
	}

	tx.Rollback()
}

func stablishConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln(err)
	}
	return db, err
}
