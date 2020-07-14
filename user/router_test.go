package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var findAll func(responseChan chan func() ([]User, error))
var findByID func(id string, responseChan chan func() (User, error))
var save func(user User, responseChan chan func() (User, error))
var delete func(id string, responseChan chan error)

type userRepositoryMock struct{}

func (u userRepositoryMock) FindAll(responseChan chan func() ([]User, error)) {
	findAll(responseChan)
}

func (u userRepositoryMock) FindByID(id string, responseChan chan func() (User, error)) {
	findByID(id, responseChan)
}

func (u userRepositoryMock) Save(user User, responseChan chan func() (User, error)) {
	save(user, responseChan)
}

func (u userRepositoryMock) Delete(id string, responseChan chan error) {
	delete(id, responseChan)
}

func TestGetEmptyUsers(t *testing.T) {
	userRepository = userRepositoryMock{}
	users := []User{}
	findAll = func(responseChan chan func() ([]User, error)) {
		responseChan <- (func() ([]User, error) { return users, nil })
	}

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
	expected := `[]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetAllUsers(t *testing.T) {
	userRepository = userRepositoryMock{}
	users := make([]User, 2)
	users[0] = User{ID: 1, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	users[1] = User{ID: 2, FirstName: "xyz", LastName: "pqr", Phone: "1234567890"}
	findAll = func(responseChan chan func() ([]User, error)) {
		responseChan <- (func() ([]User, error) { return users, nil })
	}

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
}

func TestCreateNewUser(t *testing.T) {
	userRepository = userRepositoryMock{}
	var userFromRepository = User{ID: 1, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	save = func(user User, responseChan chan func() (User, error)) {
		if user.FirstName != "Krish" || user.LastName != "Bhanushali" {
			t.Errorf("Provided user is not correct. Provided first name is %v, and last name %v", user.FirstName, user.LastName)
		}
		responseChan <- (func() (User, error) { return userFromRepository, nil })
	}

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
	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","phone":"0987654321"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestFetchUserById(t *testing.T) {
	userRepository = userRepositoryMock{}
	var userFromRepository = User{ID: 12, FirstName: "Krish", LastName: "Bhanushali", Phone: "0987654321"}
	findByID = func(id string, responseChan chan func() (User, error)) {
		if id != "12" {
			t.Errorf("Provided user id is not correct. Provided id is %v, expected %v", id, "12")
		}
		responseChan <- (func() (User, error) { return userFromRepository, nil })
	}

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
}

func TestDeleteUserById(t *testing.T) {
	userRepository = userRepositoryMock{}
	delete = func(id string, responseChan chan error) {
		if id != "22" {
			t.Errorf("Provided user id is not correct. Provided id is %v, expected %v", id, "22")
		}
		responseChan <- nil
	}

	req, err := http.NewRequest("DELETE", "/users/22", nil)
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
}
