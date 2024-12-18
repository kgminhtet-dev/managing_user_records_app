package users

type Service struct {
	repository *Repository
}

func (s *Service) GetUsers(page, limit int) ([]*User, error) {
	return nil, nil
}

func (s *Service) GetUserById(id string) (*User, error) {
	return nil, nil
}

func (s *Service) UpdateUser(id string, user *User) (*User, error) {
	return nil, nil
}

func (s *Service) DeleteUser(id string) (*User, error) {
	return nil, nil
}

func newService(repo *Repository) *Service {
	return &Service{repository: repo}
}
