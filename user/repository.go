package user

// UserRepository - asdfasd
type UserRepository interface {
	FindAll(responseChan chan func() ([]User, error))
	FindByID(id string, responseChan chan func() (User, error))
	Save(user User, responseChan chan func() (User, error))
}

// UserRepositoryStruct - asdfasd
type UserRepositoryStruct struct{}

// FindAll - asfsadf
func (r UserRepositoryStruct) FindAll(responseChan chan func() ([]User, error)) {
	FindAll(responseChan)
}

// FindByID - asfsadf
func (r UserRepositoryStruct) FindByID(id string, responseChan chan func() (User, error)) {
	FindByID(id, responseChan)
}

// Save - asfsadf
func (r UserRepositoryStruct) Save(user User, responseChan chan func() (User, error)) {
	Save(user, responseChan)
}
