package user

// Repository - asdfasd
type Repository interface {
	FindAll(responseChan chan func() ([]User, error))
	FindByID(id string, responseChan chan func() (User, error))
	Save(user User, responseChan chan func() (User, error))
	Delete(id string, responseChan chan error)
}

// RepositoryStruct - asdfasd
type RepositoryStruct struct{}

// FindAll - asfsadf
func (r RepositoryStruct) FindAll(responseChan chan func() ([]User, error)) {
	FindAll(responseChan)
}

// FindByID - asfsadf
func (r RepositoryStruct) FindByID(id string, responseChan chan func() (User, error)) {
	FindByID(id, responseChan)
}

// Save - asfsadf
func (r RepositoryStruct) Save(user User, responseChan chan func() (User, error)) {
	Save(user, responseChan)
}

// Delete - asfsadf
func (r RepositoryStruct) Delete(id string, responseChan chan error) {
	Delete(id, responseChan)
}
