package user

// UserRepository - asdfasd
type UserRepository interface {
	FindAll(responseChan chan func() ([]User, error))
}

// UserRepositoryStruct - asdfasd
type UserRepositoryStruct struct{}

// FindAll - asfsadf
func (r UserRepositoryStruct) FindAll(responseChan chan func() ([]User, error)) {
	FindAll(responseChan)
}
