package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var findAll func(responseChan chan func() ([]User, error))
var save func(user User, responseChan chan func() (User, error))

type userRepositoryMock struct{}

func (u userRepositoryMock) FindAll(responseChan chan func() ([]User, error)) {
	findAll(responseChan)
}

func (u userRepositoryMock) Save(user User, responseChan chan func() (User, error)) {
	save(user, responseChan)
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
